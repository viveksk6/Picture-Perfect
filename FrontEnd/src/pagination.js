import React, { Component } from "react";
import "./App.css";
import { connect } from "react-redux";
import { loadMovies } from "./actions/index.js";
import Pagination from "react-js-pagination";

class Paginate extends Component {
  componentDidMount = () => {
    loadMovies(1);
  };
  handlePageChange = pageNumber => {
    const { loadMovies } = this.props;
    console.log(`active page is ${pageNumber}`);
    //this.setState({ activePage: pageNumber });

    loadMovies(pageNumber);
  };

  render() {
    return (
      <div>
        <Pagination
          itemClass="page-item"
          linkClass="page-link"
          disabledClass="disabled"
          activePage={this.props.activePage}
          itemsCountPerPage={10}
          totalItemsCount={450}
          pageRangeDisplayed={5}
          onChange={this.handlePageChange}
        />
      </div>
    );
  }
}

const mapStateToProps = ({ movies, activePage }) => ({
  movies,
  activePage
});

const mapDispatchToProps = dispatch => ({
  loadMovies: pageNumber => dispatch(loadMovies(pageNumber))
});

export default connect(mapStateToProps, mapDispatchToProps)(Paginate);
