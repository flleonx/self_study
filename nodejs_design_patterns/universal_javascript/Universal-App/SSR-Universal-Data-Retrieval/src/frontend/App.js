import React from "react";
import htm from "htm";
import { Switch, Route } from "react-router-dom";
import { routes } from "./routes.js";

const html = htm.bind(React.createElement);

export class App extends React.Component {
  render() {
    return html`
      <${Switch}>
        ${routes.map(
          (routeConfig) =>
            html`<${Route} key=${routeConfig.path} ...${routeConfig} />`
        )}
      </${Switch}>
    `;
  }
}
