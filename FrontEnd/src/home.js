import React, { Component } from "react";
import Paginate from "./pagination";
import Cards from "./cards";

class Home extends Component {
  getName = () => {
    return "Now Showing";
  };
  render() {
    return (
      <div>
        <div className="left-top">
          <h2>{this.getName()}</h2>
        </div>

        <div className="left-top2">
          <Cards />
        </div>

        <div className="left-top3">
          <Paginate />
        </div>
      </div>
    );
  }
}

export default Home;
