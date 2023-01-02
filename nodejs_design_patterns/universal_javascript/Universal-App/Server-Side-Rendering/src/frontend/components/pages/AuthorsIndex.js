import React from "react";
import htm from "htm";
import { Link } from "react-router-dom";
import { Header } from "../Header.js";
import { authors } from "../../../data/authors.js";

const html = htm.bind(React.createElement);

export class AuthorsIndex extends React.Component {
  render() {
    return html`<div>
      <${Header} />
      <div>
        ${authors.map(
          (author) =>
            html`<div key=${author.id}>
          <p>
            <${Link} to="${`/author/${author.id}`}">
              ${author.name}
            </${Link}>
          </p>
        </div>`
        )}
      </div>
    </div>`;
  }
}
