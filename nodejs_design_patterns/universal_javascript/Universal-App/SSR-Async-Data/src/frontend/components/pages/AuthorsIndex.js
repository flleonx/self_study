import React from "react";
import htm from "htm";
import { Link } from "react-router-dom";
import superagent from "superagent";
import { Header } from "../Header.js";

const html = htm.bind(React.createElement);

export class AuthorsIndex extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      authors: [],
      loading: true,
    };
  }

  async componentDidMount() {
    try {
      const { body } = await superagent.get("http://localhost:3001/api/authors");
      this.setState({ loading: false, authors: body });
    } catch (error) {
      console.log(error);
      this.setState({ loading: false, authors: null });
    } 
  }

  render() {
    if (this.state.loading) {
      return html`<${Header} />
        <div>Loading...</div>`;
    }

    return html`<div>
      <${Header} />
      <div>
        ${this.state.authors.map(
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
