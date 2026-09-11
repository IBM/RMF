/*
 * Minimal live-reload plugin for the plugin dev build.
 *
 * Replaces the unmaintained `webpack-livereload-plugin` (last release 2021) with a
 * small vendored implementation, mirroring the approach `@grafana/create-plugin`
 * takes for its Rspack bundler (see `.config/rspack/liveReloadPlugin.ts` upstream).
 *
 * On each rebuild it starts (once) a WebSocket server on `port` and tells any
 * connected clients to reload. The client script is served at `/livereload.js`.
 * The matching `<script>` tag is injected into Grafana's index.html by the Docker
 * image (see `.config/Dockerfile`), so `appendScriptTag` defaults to false to
 * preserve the existing behavior.
 */
import path from 'path';
import { createServer, type Server } from 'http';
import { WebSocket, WebSocketServer } from 'ws';
import webpack, { type Compiler, type Compilation } from 'webpack';

interface LiveReloadPluginOptions {
  port?: number;
  protocol?: string;
  appendScriptTag?: boolean;
}

class LiveReloadPlugin {
  options: Required<LiveReloadPluginOptions>;
  httpServer: Server | null = null;
  server: WebSocketServer | null = null;

  constructor(options: LiveReloadPluginOptions = {}) {
    this.options = Object.assign(
      {
        port: 35729,
        protocol: 'http',
        appendScriptTag: false,
      },
      options
    );
  }

  apply(compiler: Compiler) {
    compiler.hooks.afterEmit.tap('LiveReloadPlugin', () => {
      this._startServer();
      this._notifyClient();
    });

    if (this.options.appendScriptTag) {
      compiler.hooks.thisCompilation.tap('LiveReloadPlugin', (compilation) => {
        compilation.hooks.processAssets.tap(
          {
            name: 'LiveReloadPlugin',
            stage: webpack.Compilation.PROCESS_ASSETS_STAGE_ADDITIONAL,
          },
          () => this._injectLiveReloadScript(compilation)
        );
      });
    }
  }

  _startServer() {
    if (this.server) {
      return;
    }

    const { port } = this.options;

    this.httpServer = createServer((req, res) => {
      if (req.url === '/livereload.js') {
        res.writeHead(200, { 'Content-Type': 'application/javascript' });
        res.end(this._getLiveReloadScript());
      } else {
        res.writeHead(404);
        res.end('Not Found');
      }
    });

    this.server = new WebSocketServer({ server: this.httpServer });
    this.httpServer.listen(port, () => {
      // eslint-disable-next-line no-console
      console.log(`LiveReload server started on http://localhost:${port}`);
    });
  }

  _notifyClient() {
    if (!this.server) {
      return;
    }

    this.server.clients.forEach((client) => {
      if (client.readyState === WebSocket.OPEN) {
        client.send(JSON.stringify({ action: 'reload' }));
      }
    });
  }

  _injectLiveReloadScript(compilation: Compilation) {
    const scriptTag = `<script src="http://localhost:${this.options.port}/livereload.js"></script>`;

    for (const filename of Object.keys(compilation.assets)) {
      if (path.extname(filename) !== '.html') {
        continue;
      }

      const asset = compilation.getAsset(filename);
      if (!asset) {
        continue;
      }

      const updated = asset.source.source().toString().replace('</body>', `${scriptTag}</body>`);
      compilation.updateAsset(filename, new webpack.sources.RawSource(updated));
    }
  }

  _getLiveReloadScript() {
    return `
      (function () {
        if (typeof WebSocket === 'undefined') {
          return;
        }
        const ws = new WebSocket('${this.options.protocol}://localhost:${this.options.port}');
        ws.onmessage = function (event) {
          const data = JSON.parse(event.data);
          if (data.action === 'reload') {
            window.location.reload();
          }
        };
      })();
    `;
  }
}

export default LiveReloadPlugin;
