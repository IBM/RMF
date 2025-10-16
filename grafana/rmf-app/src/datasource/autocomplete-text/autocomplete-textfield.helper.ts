/**
 * (C) Copyright IBM Corp. 2023.
 * (C) Copyright Rocket Software, Inc. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
const properties = [
  'direction',
  'boxSizing',
  'width', // on Chrome and IE, exclude the scrollbar, so the mirror div wraps exactly as the textarea does
  'height',
  'overflowX',
  'overflowY', // copy the scrollbar for IE

  'borderTopWidth',
  'borderRightWidth',
  'borderBottomWidth',
  'borderLeftWidth',
  'borderStyle',

  'paddingTop',
  'paddingRight',
  'paddingBottom',
  'paddingLeft',

  'fontStyle',
  'fontVariant',
  'fontWeight',
  'fontStretch',
  'fontSize',
  'fontSizeAdjust',
  'lineHeight',
  'fontFamily',

  'textAlign',
  'textTransform',
  'textIndent',
  'textDecoration',

  'letterSpacing',
  'wordSpacing',

  'tabSize',
  'MozTabSize',
];

const isBrowser = typeof window !== 'undefined';
const isFirefox = isBrowser && (window as any).mozInnerScreenX != null;

export const getCaretCoordinates = (element: any, position: any, options?: any) => {
  // The mirror div will replicate the textarea's style
  let div = document.createElement('div');
  div.id = 'input-autotextarea-caret-position-mirror-div';
  document.body.appendChild(div);

  let style = div.style;
  let computed = window.getComputedStyle ? window.getComputedStyle(element) : element.currentStyle; // currentStyle for IE < 9
  // let isInput = element.nodeName === 'INPUT';

  // Default textarea styles
  style.whiteSpace = 'pre-wrap';
  style.wordWrap = 'break-word'; // only for textarea-s

  // Position off-screen
  style.position = 'absolute'; // required to return coordinates properly
  style.visibility = 'hidden'; // not 'display: none' because we want rendering

  // Transfer the element's properties to the div
  properties.forEach(function (prop) {
    style[prop as any] = computed[prop];
  });

  if (isFirefox) {
    // Firefox lies about the overflow property for textareas:
    // If the string begins with "0x", the radix is 16 (hexadecimal)
    // If the string begins with "0", the radix is 8 (octal). This feature is deprecated
    // If the string begins with any other value, the radix is 10 (decimal)
    if (element.scrollHeight > parseInt(computed.height, 10)) {
      style.overflowY = 'scroll';
    }
  } else {
    style.overflow = 'hidden'; // for Chrome to not render a scrollbar; IE keeps overflowY = 'scroll'
  }

  div.textContent = element.value.substring(0, position);

  let span = document.createElement('span');
  span.textContent = element.value.substring(position) || '.'; // || because a completely empty faux span doesn't render at all
  div.appendChild(span);

  let coordinates = {
    top: span.offsetTop + parseInt(computed['borderTopWidth'], 10),
    left: span.offsetLeft + parseInt(computed['borderLeftWidth'], 10),
    height: parseInt(computed['lineHeight'], 10),
  };

  document.body.removeChild(div);

  return coordinates;
};

//Support New Grammer
export const isItFirstDotPosition = (triggerStr: string, str: string, triggerIdx: number): boolean => {
  if (triggerStr === '.' && str.toString().indexOf(triggerStr) !== triggerIdx) {
    return false;
  }
  return true;
};
