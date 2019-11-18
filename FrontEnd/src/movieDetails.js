import React, { Component } from "react";
import Button from "react-bootstrap/Button";

class movieDetails extends Component {
  render() {
    return (
      <div>
        <img src="grey.jpg" className="bg-image" />
        <img src="movie1.jpg" className="movie-image" />
        <h2 className="movie-name">Bahubali 2 - The conclusion</h2>
        <h5 className="certificate">(U/A) | Telugu | Drama</h5>
        <Button variant="primary" className="shows-button">
          View Shows
        </Button>
      </div>
    );
  }
}

export default movieDetails;
