import React from "react";
import htm from "htm";
import { Link } from "react-router-dom";

const html = htm.bind(React.createElement);

export class Header extends React.Component {
  render() {
    return html`<header>
      <h1>
        <${Link} to="/">My library</${Link}>
      </h1> 
    </header>`;
  }
}
