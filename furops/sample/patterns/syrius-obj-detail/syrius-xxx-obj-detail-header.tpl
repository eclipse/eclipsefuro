import { LitElement, html, css } from 'lit';
import { FBP } from '@furo/fbp';
import { MediaSize } from '@furo/ui5/src/lib/MediaSize.js';

import '@furo/route/src/furo-app-flow.js';
import '@furo/route/src/furo-location-updater.js';
import '@furo/route/src/furo-pages.js';

import '@furo/ui5/src/furo-ui5-card.js';
import '@furo/ui5/src/furo-ui5-header-panel.js';
import '@furo/ui5/src/furo-ui5-message-strip-display.js';
import '@furo/ui5/src/furo-ui5-message-strip.js';
import '@furo/ui5/src/furo-ui5-typerenderer-labeled.js';
import '@furo/ui5/src/furo-ui5-dialog.js';

import '@ui5/webcomponents/dist/TabContainer.js';
import '@ui5/webcomponents/dist/Tab.js';
import '@ui5/webcomponents/dist/Button.js';
import '@ui5/webcomponents/dist/Popover.js';

import '@ui5/webcomponents-fiori/dist/ShellBar.js';
import '@ui5/webcomponents-fiori/dist/ShellBarItem.js';

/**
 * `view-syrius-xxx-obj-detail`
 * Object Details is a view that gives you a complete set of data tied to particular object.
 *
 * @customElement
 * @appliesMixin FBP
 */
class ViewSyrius{{Name}}ObjDetail extends FBP(LitElement) {
  constructor() {
    super();
    this._currentTab = 'basics';
  }

  /**
   * flow is ready lifecycle method
   */
  _FBPReady() {
    super._FBPReady();
    // this._FBPTraceWires();

    /**
     * Responsive spacing system
     */
    window.addEventListener(
      'resize',
      MediaSize.DebounceBuilder(() => {
        this.setAttribute('media-size', MediaSize.GetMediaSize());
      }, MediaSize.HANDLE_RESIZE_DEBOUNCE_RATE),
    );
    // initial size
    this.setAttribute('media-size', MediaSize.GetMediaSize());

    /**
     * Tab controller
     * Ensures that the correct tab is activated in the tab container after a page reload.
     */
    this._FBPAddWireHook('--pageHashChanged', e => {
      this._currentTab = e.hash.tab;
      if (this._currentTab === undefined) {
        // set default
        this._currentTab = 'basics';
      }
      this.requestUpdate();
    });
  }

  /**
   * checks if the current tab is selected
   * @param tab
   * @return {boolean}
   * @private
   */
  _checkInitialTab(tab) {
    return this._currentTab === tab;
  }

  /**
   * CSS Styles
   * @private
   * @return {CSSResult}
   */
  static get styles() {
    // language=CSS
    return css`
      /* width */
      ::-webkit-scrollbar {
        width: var(--sapScrollBar_Dimension, 0.75rem);
      }

      /* Track */
      ::-webkit-scrollbar-track {
        background: var(--sapScrollBar_TrackColor, #090b0d);
      }

      /* Handle */
      ::-webkit-scrollbar-thumb {
        background: var(--sapScrollBar_FaceColor, #91c8f6);
      }

      /* Handle on hover */
      ::-webkit-scrollbar-thumb:hover {
        background: var(--sapScrollBar_Hover_FaceColor, #4a5a6a);
      }

      :host {
        display: block;
        height: 100%;
      }

      :host([hidden]) {
        display: none;
      }

      :host([media-size='XXL']) .section {
        padding: 2rem 3rem 1rem 3rem;
      }

      :host([media-size='XL']) .section {
        padding: 2rem 3rem 1rem 3rem;
      }

      :host([media-size='L']) .section {
        padding: 1rem 2rem 0 2rem;
      }

      :host([media-size='M']) .section {
        padding: 0.625rem 2rem 0 2rem;
      }

      :host([media-size='S']) .section {
        padding: 0.625rem 1rem 0 1rem;
      }
    `;
  }

  /**
   * Template render function
   * @returns {TemplateResult}
   * @private
   */
  render() {
    // language=HTML
    return html`
      <furo-vertical-flex>
        <!-- Shell Bar with standard functionality -->
        <ui5-shellbar
          primary-title="$SHELL-BAR-TITLE$"
          secondary-title="$SHELL-BAR-SECONDARY-TITLE$"
          @-profile-click="--userRequested(*.detail.targetRef)"
        >
          <ui5-avatar slot="profile" initials=""></ui5-avatar>
          <ui5-button
            icon="nav-back"
            slot="startButton"
            @-click="^^return-to-last-waypoint"
          ></ui5-button>
          <ui5-shellbar-item
            id="techView"
            icon="tnt/values"
            title="All attributes"
            text="All Attributes"
            @-click="--openAllAttributes"
          ></ui5-shellbar-item>
          <ui5-shellbar-item
            id="history"
            icon="history"
            text="History"
            title="History"
          ></ui5-shellbar-item>
          <ui5-shellbar-item id="help" icon="sys-help" text="Help" title="Help"></ui5-shellbar-item>
        </ui5-shellbar>

        <!-- shows the user profile related menu -->
        <ui5-popover ƒ-show-at="--userRequested" placement-type="Bottom">
          <div class="popover-header">
            <!-- Uncomment the line below to add a title -->
            <!-- <ui5-title style="padding: 0.25rem 1rem 0rem 1rem"></ui5-title> -->
          </div>

          <div class="popover-content">
            <ui5-list separators="None">
              <ui5-li icon="settings">Language</ui5-li>
              <ui5-li icon="sys-help">Help</ui5-li>
            </ui5-list>
          </div>
        </ui5-popover>

        <!-- dialog with the all attributes view -->
        <furo-ui5-dialog
          ƒ-show="--openAllAttributes"
          ƒ-close="--closeDialogClicked"
          header-text=""
          stretch="true"
        >
          <!-- Add your all attribute view according the sample below -->
          <!-- <view-xxx-all-attributes ƒ-bind-data="--dao"></view-xxx-all-attributes> -->

          <div slot="footer">
            <furo-ui5-button @-click="--closeDialogClicked">Close </furo-ui5-button>
          </div>
        </furo-ui5-dialog>

        <!-- Business object header, shows the most important attributes -->
        <furo-ui5-header-panel ƒ-bind-header-text="" ƒ-bind-secondary-text="">
          <furo-horizontal-flex space slot="action">
            <!-- you can add action controls -->
            <!-- <furo-ui5-button>my action</furo-ui5-button> -->
          </furo-horizontal-flex>

          <furo-horizontal-flex>
            <furo-form-layouter six flex>
              <div>
                <!-- column 1-->
                <!-- <furo-ui5-typerenderer-labeled ƒ-bind-data="--dao(*.display_name)"></furo-ui5-typerenderer-labeled> -->
              </div>
              <div>
                <!-- column 2-->
              </div>
              <div>
                <!-- column 3-->
              </div>
              <div>
                <!-- column 4-->
              </div>
              <div>
                <!-- column 5-->
              </div>
              <div>
                <!-- column 6-->
              </div>
            </furo-form-layouter>
          </furo-horizontal-flex>
        </furo-ui5-header-panel>

        <!-- Tab navigation bar for a better structure (optional) -->
        <ui5-tabcontainer class="full-width" collapsed fixed @-tab-select="--subTabSelected">
          <!-- The ui5-tab represents a selectable item inside an ui5-tabcontainer. -->
          <ui5-tab
            text="Basic"
            data-tab="basics"
            ?selected="${this._checkInitialTab('basics')}"
          ></ui5-tab>
        </ui5-tabcontainer>

        <furo-location-updater ƒ-set-hash="--subTabSelected(*.tab.dataset)"></furo-location-updater>

        <!-- gRPC localized messages -->
        <furo-ui5-message-strip-display class="section"></furo-ui5-message-strip-display>
        <furo-ui5-message-strip
          message="SYR-999, General Error"
          ƒ-show-error="--notImplemented, --badGateway"
          ƒ-show-grpc-localized-message="--grpcError"
        ></furo-ui5-message-strip>

        <!-- Container that holds the different panels -->
        <furo-pages scroll ƒ-activate-page="--pageHashChanged(*.hash.tab)" default="basics">
          <div name="basics"></div>
        </furo-pages>
      </furo-vertical-flex>

      <!-- Data model and backend communication component -->
      <!-- Entity-Resolver-Component -->

      <furo-document-title prefix="Detail: " ƒ-bind-title=""></furo-document-title>
    `;
  }
}

window.customElements.define('view-syrius-xxx-obj-detail', ViewSyriusXxxObjDetail);
