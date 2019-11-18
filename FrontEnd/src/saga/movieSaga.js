import { put, call, takeEvery } from "redux-saga/effects";

import { setMovies, setError } from "../actions";
import { MOVIES } from "../constants";
import { fetchMovies } from "../api/index.js";

//const page;

export function* handleMoviesLoad() {
  try {
    //console.log("1");
    const movies = yield call(fetchMovies);
    //console.log(page);
    yield put(setMovies(movies));
  } catch (error) {
    yield put(setError(error.toString()));
  }
}

export default function* watchMoviesLoad() {
  const { page } = yield takeEvery(MOVIES.LOAD, handleMoviesLoad);
  console.log(page);
}
