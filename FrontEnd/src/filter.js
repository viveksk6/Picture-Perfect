import React, { Component } from "react";
import Dropdown from "react-bootstrap/Dropdown";
import DropdownButton from "react-bootstrap/DropdownButton";

class Filter extends Component {
  render = () => {
    return (
      <DropdownButton
        className="left-top"
        id="dropdown-item-button"
        title="Filter"
      >
        <DropdownButton id="drop2" title="Language">
          <Dropdown.Item as="button">English</Dropdown.Item>
          <Dropdown.Item as="button">French</Dropdown.Item>
          <Dropdown.Item as="button">German</Dropdown.Item>
        </DropdownButton>
        <DropdownButton id="drop3" title="Genre">
          <Dropdown.Item as="button">Horror</Dropdown.Item>
          <Dropdown.Item as="button">Comedy</Dropdown.Item>
          <Dropdown.Item as="button">Romance</Dropdown.Item>
          <Dropdown.Item as="button">Thriller</Dropdown.Item>
          <Dropdown.Item as="button">Drama</Dropdown.Item>
        </DropdownButton>
      </DropdownButton>
    );
  };
}

export default Filter;
