import React, { Component } from "react";
import Pagination from "react-bootstrap/Pagination";

class Paginate extends Component {
  addItem = () => {
    let active = 2;
    let items = [];
    for (let number = 1; number <= 5; number++) {
      items.push(
        <Pagination.Item key={number} active={number === active}>
          {number}
        </Pagination.Item>
      );
    }
    return items;
  };
  render() {
    return (
      <div>
        <Pagination>{this.addItem()}</Pagination>
      </div>
    );
  }
}

export default Paginate;
