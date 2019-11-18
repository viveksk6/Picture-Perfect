import { MOVIES } from "../constants";

const pageReducer = (state = 1, action) => {
  switch (action.type) {
    case MOVIES.LOAD_SUCCESS:
      return [...state, ...action.pageNumber];
    default:
      return state;
  }
};

export default pageReducer;
