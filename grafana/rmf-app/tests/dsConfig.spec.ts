import { test, expect } from '@playwright/test';

const LEGACY_DATA_SOURCES = [
  {
    uid: 'no-ssl-no-auth-legacy',
    url: 'http://dds.test',
    basicAuth: false,
    basicAuthUser: '',
    jsonData: {
      timeout: '60',
      tlsSkipVerify: false,
    },
  },
  {
    uid: 'ssl-and-auth-legacy',
    url: 'https://dds.test:8803',
    basicAuth: true,
    basicAuthUser: 'tsxxx',
    jsonData: {
      timeout: '60',
      tlsSkipVerify: true,
    },
  },
];

test.describe('Legacy format conversion', () => {
  for (let dataSource of LEGACY_DATA_SOURCES) {
    test(`${dataSource.uid}`, async ({ page }) => {
      await page.goto(`/connections/datasources/edit/${dataSource.uid}`);

      const urlField = page.locator('_react=[label = "DDS URL"]');
      await expect(urlField).toHaveText('DDS URL');
      await expect(urlField.getByRole('textbox')).toHaveValue(dataSource.url);

      const timeoutField = page.locator('_react=[label = "Timeout"]');
      await expect(timeoutField).toHaveText('Timeout');
      await expect(timeoutField.getByRole('textbox')).toHaveValue(dataSource.jsonData.timeout);

      const baseAuthField = page.locator('_react=[label = "Basic Auth"]');
      await expect(baseAuthField).toHaveText('Basic Auth');
      await expect(baseAuthField.getByRole('checkbox')).toBeChecked({ checked: dataSource.basicAuth });

      const tlsVerifyField = page.locator('_react=[label = "Skip TLS Verify"]');
      await expect(tlsVerifyField).toHaveText('Skip TLS Verify');
      await expect(tlsVerifyField.getByRole('checkbox')).toBeChecked({ checked: dataSource.jsonData.tlsSkipVerify });

      const userField = page.locator('_react=[label = "User"]');
      const passwordField = page.locator('_react=[label = "Password"]');
      if (dataSource.basicAuth) {
        await expect(userField).toHaveText('User');
        await expect(userField.getByRole('textbox')).toHaveValue(dataSource.basicAuthUser);
        await expect(passwordField).toHaveText('Password');
        await expect(passwordField.getByRole('textbox')).toHaveValue('');
      } else {
        await expect(userField).not.toBeVisible();
        await expect(passwordField).not.toBeVisible();
      }
    });
  }
});
