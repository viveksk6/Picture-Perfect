import { MOVIES } from "../constants";

const loadMovies = () => ({
  type: MOVIES.LOAD
});

const setMovies = movies => ({
  type: MOVIES.LOAD_SUCCESS,
  movies
  //pageNumber
});

const setError = error => ({
  type: MOVIES.LOAD_FAIL,
  error
});

export { loadMovies, setMovies, setError };
