import React from "react";
import ReactDOM from "react-dom";

const h = React.createElement;

class Hello extends React.Component {
  render() {
    return h("h1", null, ["Hello ", this.props.name || "World"]);
  }
}

ReactDOM.render(
  h(Hello, { name: "React" }),
  document.getElementsByTagName("body")[0]
);
