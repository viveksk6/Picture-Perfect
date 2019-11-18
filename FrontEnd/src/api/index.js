//import movieData from "../movieData.json";

const pageSize = `3`;
const pageNo = `1`;
const URL = `http://localhost:10000/all`;
//const cor = `https://cors-anywhere.herokuapp.com/`;
//const page = `all`;

const fetchMovies = async () => {
  const response = await fetch(`${URL}?pageSize=${pageSize}&pageNo=${pageNo}`);
  const data = await response.json();
  if (response.status >= 400) {
    throw new Error(data.Error);
  }
  console.log(data);
  return data;
  //return movieData;
};

export { fetchMovies };
