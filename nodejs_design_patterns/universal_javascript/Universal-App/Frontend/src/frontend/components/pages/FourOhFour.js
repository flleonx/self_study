import React from "react";
import htm from "htm";
import { Link } from "react-router-dom";
import { Header } from "../Header.js";

const html = htm.bind(React.createElement);

export class FourOhFour extends React.Component {
  render() {
    return html`<div>
      <${Header} />
      <div>
        <h2>404</h2>
        <h3>${this.props.error || "Page not found"}</h3>
        <${Link} to="/">Go back to the home page</${Link}>
      </div>
    </div>`;
  }
}
