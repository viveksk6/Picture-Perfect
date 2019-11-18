import React, { Component } from "react";
import { Provider } from "react-redux";
import "bootstrap/dist/css/bootstrap.min.css";
import Navigation from "./navigation";
import Home from "./home";
import configureStore from "./store";
import movieDetails from "./movieDetails";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import "./App.css";

const store = configureStore();

class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <div className="App">
          <div className="left-top0">
            <Navigation />
          </div>
          <div>
            <Router>
              <Link to="/"></Link>
              <Switch>
                <Route exact path="/">
                  <Home />
                </Route>
                <Route path="/:movieID" component={movieDetails} exact />
              </Switch>
            </Router>
          </div>
        </div>
      </Provider>
    );
  }
}

export default App;
