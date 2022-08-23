// TODO : 
// Search
// Fav list | almost...
// Affichage

import { useEffect, useState } from "react";
import MovieCard from "../components/MovieCard";
import movieService from "../services/movie.service";
import Modal from "../components/Modal";
import ListShow from "../components/ListShow";

export default function Home() {

  const [discoveryMovies, setDiscoveryMovies] = useState();
  const [popularMovies, setPopularMovies] = useState();
  const [genres, setGenres] = useState();

  // showModal show the full movie informations available
  const [showModal, setShowModal] = useState(false);

  // showModal show the full movie informations available
  const [modalMovie, setModalMovie] = useState({ title: "", overview: "", release_date: "" });

  useEffect(() => {
    movieService.getDiscoveryMovies().then((data) => {
      setDiscoveryMovies(data.results);
    });

    movieService.getPopularMovies().then((data) => {
      setPopularMovies(data.results);
      console.log(data.results[0]);
    });

    movieService.getGenres().then((data) => {
      setGenres(data.genres);
    })
  }, []);

  return (
    <>
      <Modal title={modalMovie.title} isActive={showModal} closeFunction={() => setShowModal(!showModal)} type="information">
        <p>{modalMovie.release_date.split("-")[0]}</p>
        <p>{modalMovie.overview}</p>
      </Modal>

      <ListShow listName="Discovery" movies={discoveryMovies} genres={genres} pshowModal={showModal} psetShowModal={setShowModal} psetModalMovie={setModalMovie}/>



      <br />
      <hr />
      <br />

      <h1>Popular</h1>
      <div className="movies__grid">
        {popularMovies && popularMovies.map((currrentMovie) => (
          <MovieCard movie={currrentMovie} key={currrentMovie.id} genres={genres} pshowModal={showModal} psetShowModal={setShowModal} psetModalMovie={setModalMovie} />
        ))}

      </div>
    </>
  )
}
