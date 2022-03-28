import { LitElement, html, css } from 'lit';
import {FBP} from "@furo/fbp";

/**
 * `{{.Var.ElementName}}`
 *
 * todo Add a extended description or delete this line
 *
 * @summary {{.Var.Description}}
 * @customElement {{.Var.ElementName}}
 * @appliesMixin FBP
 */
class {{.Var.ClassName}} extends FBP(LitElement) {


  {{if .Var.Reactive}}
  /**
   * @private
   * @return {Object}
   */
  static get properties() {
    return {
    };
  }

  {{end}}
  /**
   * flow is ready lifecycle method
   */
  _FBPReady(){
    super._FBPReady();
    // this._FBPTraceWires()
  }

  /**
   * Themable Styles
   * @private
   * @return {CSSResult}
   */
  static get styles() {
    // language=CSS
    return css`
        :host {
            display: block;
        }

        :host([hidden]) {
            display: none;
        }
    `
  }


  /**
   * @private
   * @returns {TemplateResult}
   * @private
   */
  render() {
    // language=HTML
    return html`
      <p>Hej, welcome</p>
    `;
  }
}

window.customElements.define('{{.Var.ElementName}}', {{.Var.ClassName}});
