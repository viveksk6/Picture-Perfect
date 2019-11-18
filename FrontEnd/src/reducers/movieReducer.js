import { MOVIES } from "../constants";
import movieData from "../movieData.json";

const initialState = {
  movies: movieData
};

const movieReducer = (state = initialState, action) => {
  if (action.type === MOVIES.LOAD_SUCCESS) {
    return { movies: action.movies };
  }
  return state;
};

export default movieReducer;
