import React, { Component } from "react";
import Navbar from "react-bootstrap/Navbar";
import Form from "react-bootstrap/Form";
import FormControl from "react-bootstrap/FormControl";
import Button from "react-bootstrap/Button";

class Navigation extends Component {
  render = () => {
    return (
      <Navbar bg="dark" variant="dark" fixed="top">
        <Form inline>
          <FormControl type="text" placeholder="Search" className="mr-sm-2" />
          <Button variant="outline-info">Search</Button>
        </Form>

        <Navbar.Collapse className="justify-content-end">
          <Navbar.Text>Hello Vivek</Navbar.Text>
        </Navbar.Collapse>
      </Navbar>
    );
  };
}

export default Navigation;
