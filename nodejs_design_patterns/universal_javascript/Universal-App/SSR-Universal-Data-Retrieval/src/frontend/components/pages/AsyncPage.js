import React from "react";

export class AsyncPage extends React.Component {
  static async preloadAsyncData(props) {
    throw new Error("Must be implemented by sub class");
  }

  render() {
    throw new Error("Must be implemented by sub class");
  }

  constructor(props) {
    super(props);
    const location = props.match.url;
    this.hasData = false;
    let staticData;
    let staticError;
    const staticContext =
      typeof window !== "undefined"
        ? window.__STATIC_CONTEXT__ // client-side
        : this.props.staticContext; // server-side

    if (staticContext && staticContext[location]) {
      const { data, err } = staticContext[location];
      staticData = data;
      staticError = err;
      this.hasStaticData = true;

      typeof window !== "undefined" && delete staticContext[location];
    }

    this.state = {
      ...staticData,
      staticError,
      loading: !this.hasStaticData,
    };
  }

  async componentDidMount() {
    if (!this.hasStaticData) {
      let staticData;
      let staticError;

      try {
        const data = await this.constructor.preloadAsyncData(this.props);
        staticData = data;
      } catch (error) {
        staticError = error;
      }

      this.setState({
        ...staticData,
        loading: false,
        staticError,
      });
    }
  }
}
