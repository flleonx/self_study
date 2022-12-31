import React from "react";
import ReactDOM from "react-dom";
import htm from "htm";

const html = htm.bind(React.createElement);

class Hello extends React.Component {
  render() {
    return html`<h1>Hello ${this.props.name} || 'World'</h1>`;
  }
}

ReactDOM.render(
  html`<${Hello} name="React" />`,
  document.getElementsByTagName("body")[0]
);
