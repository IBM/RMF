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
import { PropTypes } from '../common/types';
import React, { createRef, PureComponent } from 'react';
// import getCaretCoordinates from 'textarea-caret';
import './autocomplete-textfield.css';
import { getCaretCoordinates, isItFirstDotPosition } from './autocomplete-textfield.helper';

const KEY_UP = 38;
const KEY_DOWN = 40;
const KEY_RETURN = 13;
const KEY_ENTER = 14;
const KEY_ESCAPE = 27;
const KEY_TAB = 9;

const OPTION_LIST_Y_OFFSET = 10;
const OPTION_LIST_MIN_WIDTH = 100;

type Props = PropTypes;
type state = PropTypes;

export class AutocompleteTextField extends PureComponent<Props, state> {
  // class AutocompleteTextField extends React.Component {

  recentValue: string;
  enableSpaceRemovers: boolean;
  refInput: any;

  constructor(Props: any) {
    super(Props);

    this.isTrigger = this.isTrigger.bind(this);
    this.arrayTriggerMatch = this.arrayTriggerMatch.bind(this);
    this.getMatch = this.getMatch.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleKeyDown = this.handleKeyDown.bind(this);
    this.handleKeyUp = this.handleKeyUp.bind(this);
    this.handleMouseUp = this.handleMouseUp.bind(this);
    this.handleResize = this.handleResize.bind(this);
    this.handleSelection = this.handleSelection.bind(this);
    this.updateCaretPosition = this.updateCaretPosition.bind(this);
    this.updateHelper = this.updateHelper.bind(this);
    this.resetHelper = this.resetHelper.bind(this);
    this.renderAutocompleteList = this.renderAutocompleteList.bind(this);
    this.onBlur = this.onBlur.bind(this);

    this.state = { ...this.props };

    this.recentValue = Props.defaultValue;
    this.enableSpaceRemovers = false;
    this.refInput = createRef();
  }

  componentDidMount() {
    window.addEventListener('resize', this.handleResize);
  }

  componentDidUpdate(prevProps: any) {
    const { options } = this.props;
    const { caret } = this.state;

    if (options.length !== prevProps.options.length) {
      this.updateHelper(this.recentValue, caret, options);
    }
  }

  componentWillUnmount() {
    window.removeEventListener('resize', this.handleResize);
  }

  getMatch(str: any, caret: any, providedOptions: any) {
    const { trigger, matchAny, regex } = this.props;
    const re = new RegExp(regex as any);

    let triggers = trigger;
    if (!Array.isArray(triggers)) {
      triggers = new Array(trigger);
    }
    triggers.sort();

    const providedOptionsObject = providedOptions;
    if (Array.isArray(providedOptions)) {
      triggers.forEach((triggerStr: any) => {
        providedOptionsObject[triggerStr] = providedOptions;
      });
    }

    const triggersMatch = this.arrayTriggerMatch(triggers, re);
    let slugData = null;

    for (let triggersIndex = 0; triggersIndex < triggersMatch.length; triggersIndex++) {
      const { triggerStr, triggerMatch, triggerLength } = triggersMatch[triggersIndex];

      for (let i = caret - 1; i >= 0; --i) {
        const substr = str.substring(i, caret);
        const match = substr.match(re);
        let matchstart = -1;

        if (triggerLength > 0) {
          const triggerIdx = triggerMatch ? i : i - triggerLength + 1;

          if (triggerIdx < 0) {
            // out of input
            break;
          }

          //Support New Grammer
          if (this.isTrigger(triggerStr, str, triggerIdx) && isItFirstDotPosition(triggerStr, str, triggerIdx)) {
            matchstart = triggerIdx + triggerLength;
          }

          if (!match && matchstart < 0) {
            break;
          }
        } else {
          if (match && i > 0) {
            // find first non-matching character or begin of input
            continue;
          }
          matchstart = i === 0 && match ? 0 : i + 1;

          if (caret - matchstart === 0) {
            // matched slug is empty
            break;
          }
        }

        if (matchstart >= 0) {
          const triggerOptions = providedOptionsObject[triggerStr];
          if (triggerOptions == null) {
            continue;
          }

          const matchedSlug = str.substring(matchstart, caret);

          const options = triggerOptions.filter((slug: any) => {
            const idx = slug.toLowerCase().indexOf(matchedSlug.toLowerCase());
            return idx !== -1 && (matchAny || idx === 0);
          });

          const currTrigger = triggerStr;
          const matchlength = matchedSlug.length;

          if (slugData === null) {
            slugData = {
              trigger: currTrigger,
              matchstart,
              matchlength,
              options,
            };
          } else {
            slugData = {
              ...(slugData as any),
              trigger: currTrigger,
              matchstart,
              matchlength,
              options,
            };
          }
        }
      }
    }

    return slugData;
  }

  arrayTriggerMatch(triggers: any, re: any) {
    const triggersMatch = triggers.map((trigger: any) => ({
      triggerStr: trigger,
      triggerMatch: trigger.match(re),
      triggerLength: trigger.length,
    }));

    return triggersMatch;
  }

  isTrigger(trigger: any, str: any, i: any) {
    if (!trigger || !trigger.length) {
      return true;
    }

    if (str.substr(i, trigger.length) === trigger) {
      return true;
    }

    return false;
  }

  handleChange(e: any) {
    const {
      // onChange,
      handleOnTextChange,
      options,
      spaceremovers,
      spacer,
      value,
    } = this.props;

    const old = this.recentValue;
    const str = e.target.value;
    const caret = e.target.selectionEnd;
    // e.target.selectionStart; e.target.selectionEnd;
    // const caret = getInputSelection(e.target).end;

    if (!str.length) {
      this.setState({ helpervisible: false });
    }

    this.recentValue = str;

    this.setState({ caret, value: e.target.value });

    if (!str.length || !caret) {
      return handleOnTextChange(e.target.value, e);
      // CHG return onChange (e.target.value);
    }

    // '@wonderjenny ,|' -> '@wonderjenny, |'
    if (
      this.enableSpaceRemovers &&
      (this.state.spaceremovers as any).length &&
      str.length > 2 &&
      (spacer as any).length
    ) {
      for (let i = 0; i < Math.max(old.length, str.length); ++i) {
        if (old[i] !== str[i]) {
          if (
            i >= 2 &&
            str[i - 1] === spacer &&
            (spaceremovers as any).indexOf(str[i - 2]) === -1 &&
            (spaceremovers as any).indexOf(str[i]) !== -1 &&
            this.getMatch(str.substring(0, i - 2), caret - 3, options)
          ) {
            const newValue = `${str.slice(0, i - 1)}${str.slice(i, i + 1)}${str.slice(i - 1, i)}${str.slice(i + 1)}`;

            this.updateCaretPosition(i + 1);
            this.refInput.current.value = newValue;

            if (!value) {
              this.setState({ value: newValue });
            }
            return handleOnTextChange(newValue, e);
            // CHG return onChange (newValue);
          }

          break;
        }
      }

      this.enableSpaceRemovers = false;
    }

    this.updateHelper(str, caret, options);

    if (!value) {
      this.setState({ value: e.target.value });
    }
    return handleOnTextChange(e.target.value, e);
    // CHG return onChange(e.target.value);
  }

  handleKeyDown(event: any) {
    const { handleOnKeyDown, helpervisible, options, selection } = this.state;
    const { passthroughenter } = this.props;

    handleOnKeyDown(event);
    if (helpervisible) {
      switch (event.keyCode) {
        case KEY_ESCAPE:
          event.preventDefault();
          this.resetHelper();
          break;
        case KEY_UP:
          event.preventDefault();
          this.setState({ selection: (options.length + selection - 1) % options.length });
          break;
        case KEY_DOWN:
          event.preventDefault();
          this.setState({ selection: (selection + 1) % options.length });
          break;
        case KEY_ENTER:
        case KEY_RETURN:
          if (!passthroughenter) {
            event.preventDefault();
          }
          this.handleSelection(selection);
          break;
        case KEY_TAB:
          this.handleSelection(selection);
          break;
        default:
          break;
      }
    }
  }

  handleKeyUp(event: any) {
    // const { handleOnKeyDown } = this.props;
    // if (event.key === 'Backspace' || event.key === 'Delete') {
    //   handleOnKeyDown(event);
    // }
  }

  handleMouseUp(event: any) {
    const { helpervisible } = this.state;
    // const { defaultValue } = this.props;

    if (!helpervisible) {
      // CHG onKeyDown(event);
      // if (defaultValue !== undefined && defaultValue?.length > 0 && defaultValue?.length > event.target.selectionEnd) {
      //   event.preventDefault();
      //   event.target.selectionStart = event.target.selectionEnd = defaultValue?.length + 1;
      // }
    }
  }

  onBlur(event: any) {
    if (this !== undefined && this.props !== undefined) {
      const { onBlur } = this.props;
      if (onBlur && (document.getElementsByClassName("autocomplete-input-list").length === 0)) {
        onBlur(event);
      }
    }
  }

  handleResize() {
    this.setState({ helpervisible: false });
  }

  handleSelection(idx: any, matchSucess = false) {
    const { spacer, handleOnDropDownItemSelect } = this.props;
    const { matchstart, matchlength, options, trigger } = this.state;

    const slug = options[idx];
    const value = this.recentValue;
    const part1 = trigger.length === 0 ? '' : value.substring(0, (matchstart as any) - trigger.length);
    const part2 = value.substring((matchstart as any) + (matchSucess ? (matchstart as any) - 1 : matchlength));

    const event = { target: this.refInput.current };
    const changedStr = trigger + slug; // CHG changeOnSelect(trigger, slug);

    // Text area matching string entered
    const changedStrLocal = changedStr.replace('@', '').replace('=', '');
    if (changedStrLocal.trim().toUpperCase() === part2.trim().toUpperCase()) {
      event.target.value = `${part1}${changedStr}${spacer}`;
    } else {
      event.target.value = `${part1}${changedStr}${spacer}${part2}`;
    }
    this.handleChange(event);
    handleOnDropDownItemSelect(event.target.value);

    this.resetHelper();

    this.updateCaretPosition(part1.length + changedStr.length + 1);

    this.enableSpaceRemovers = true;
  }

  updateCaretPosition(caret: any) {
    /// CHG
    // this.setState({ caret }, () => setCaretPosition(this.refInput.current, caret));
    this.refInput.current.focus();
    this.refInput.current.selectionEnd = caret;
    this.setState({ caret });
  }

  updateHelper(str: any, caret: any, options: any) {
    const input = this.refInput.current;

    const slug = this.getMatch(str, caret, options);

    if (slug) {
      const caretPos = getCaretCoordinates(input, caret);
      const rect = input.getBoundingClientRect();

      const top = caretPos.top + input.offsetTop;
      const left = Math.min(
        caretPos.left + input.offsetLeft - OPTION_LIST_Y_OFFSET,
        input.offsetLeft + rect.width - OPTION_LIST_MIN_WIDTH
      );

      const { minchars, handleOnRequestOptions, requestonlyifnooptions } = this.props;
      if (
        slug.matchlength >= (minchars as any) &&
        (slug.options.length > 1 || (slug.options.length === 1 && slug.options[0].length !== slug.matchlength))
      ) {
        this.setState({
          helpervisible: true,
          top,
          left,
          ...slug,
        });
      } else {
        if (!requestonlyifnooptions || !slug.options.length) {
          handleOnRequestOptions(str.substr(slug.matchstart, slug.matchlength));
        }

        const { helpervisible } = this.state;
        if (
          helpervisible &&
          slug.matchlength > 0 &&
          slug.options.length === 1 &&
          slug.matchlength === slug.options[0].trim().length
        ) {
          //Find correct matching item index in thelist
          // let idxList = options[slug.trigger];
          const { options: idxList } = this.state;
          let idx = 0;
          if (idxList && idxList.length > 0) {
            idx = (idxList as []).findIndex((val: string) => val.toUpperCase() === slug.options[0].toUpperCase());
            idx = idx < 0 ? 0 : idx;
          }
          this.handleSelection(idx, true);
        }
        this.resetHelper();
      }
    } else {
      this.resetHelper();
    }
  }

  resetHelper() {
    this.setState({ helpervisible: false, selection: 0 });
  }

  renderAutocompleteList() {
    const { helpervisible, left, matchstart, matchlength, options, selection, top, value } = this.state;

    if (!helpervisible) {
      return null;
    }

    const { maxOptions, offsetx, offsety } = this.props;

    if (options.length === 0) {
      return null;
    }

    if (selection >= options.length) {
      this.setState({ selection: 0 });

      return null;
    }

    const optionNumber = maxOptions === 0 ? options.length : maxOptions;

    const helperOptions = options.slice(0, optionNumber + 2).map((val: any, idx: number) => {
      const highlightStart = val
        .toLowerCase()
        .indexOf(value?.substring(matchstart as any, (matchstart as any) + matchlength).toLowerCase());

      return (
        <li
          className={idx === selection ? 'active' : undefined}
          key={val}
          onClick={() => {
            this.handleSelection(idx);
          }}
          onMouseEnter={() => {
            this.setState({ selection: idx });
          }}
        >
          {val.slice(0, highlightStart)}
          <strong>{val.substr(highlightStart, matchlength)}</strong>
          {val.slice(highlightStart + matchlength)}
        </li>
      );
    });
    helperOptions.tabindex = 0;
    return (
      <>
        <ul
          className="autocomplete-input-list"
          style={{ left: (left as any) + offsetx, top: (top as any) + offsety + 20 }}
        >
          {/* TBD: Need to Remove all popup search related code */}
          {/* <input
            type="text"
            onKeyUp={this.filterAutoPopupList}
            className="search-autocomplete-input"
            placeholder="Search for names.."
            style={{ left: (left as any) + offsetx, display: (helperOptions as any).length > 6 ? 'block' : 'none' }}
          ></input> */}
          {helperOptions}
        </ul>
      </>
    );
  }

  filterAutoPopupList() {
    // Declare variables
    let input, filter, ul, li, txtValue;
    input = document.getElementsByClassName('search-autocomplete-input');
    if (input !== null && input !== undefined) {
      filter = (input[0] as any).value.toUpperCase();
      ul = document.getElementsByClassName('autocomplete-input-list');
      li = (ul[0] as any).getElementsByTagName('li');

      // Loop through all list items, and hide those who don't match the search query
      for (let rowIndex = 0; rowIndex < li.length; rowIndex++) {
        txtValue = li[rowIndex].textContent || li[rowIndex].innerText;
        if (txtValue.toUpperCase().indexOf(filter) > -1) {
          li[rowIndex].style.display = '';
        } else {
          li[rowIndex].style.display = 'none';
        }
      }
    }
  }

  render() {
    const { Component, defaultValue, disabled, onBlur, value, ...rest } = this.props;

    const { value: stateValue } = this.state;

    const propagated = Object.assign({}, rest);
    /// CHG PropTypes=>autoCompleteDefaultProps
    // Object.keys(autoCompleteDefaultProps).forEach((k: any) => { delete propagated[k] });

    let val = '';

    if (typeof stateValue !== 'undefined' && stateValue !== null) {
      val = stateValue;
    } else if (typeof value !== 'undefined' && value !== null) {
      val = value;
    } else if (defaultValue) {
      val = defaultValue;
    }

    return (
      <>
        <Component
          className="component-autocomplete"
          disabled={disabled}
          ref={this.refInput}
          onChange={this.handleChange}
          onKeyDown={this.handleKeyDown}
          onKeyUp={this.handleKeyUp}
          onMouseUp={this.handleMouseUp}
          value={val}
          onBlur={this.onBlur}
          // onCut={(e: any) => {
          //   e.preventDefault();
          //   return false;
          // }}
          // onPaste={(e: any) => {
          //   e.preventDefault();
          //   return false;
          // }}
          {...propagated}
        />
        {this.renderAutocompleteList()}
      </>
    );
  }
}
