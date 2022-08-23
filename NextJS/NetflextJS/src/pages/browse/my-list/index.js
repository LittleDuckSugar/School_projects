import { useEffect, useState } from "react";
import movieService from "../../../services/movie.service";
import MovieCard from "../../../components/MovieCard";
import Modal from "../../../components/Modal";


const Index = () => {
    const [movieList, setMovieList] = useState([]);

    const [genres, setGenres] = useState();

    const [showModal, setShowModal] = useState(false);
    const [modalMovie, setModalMovie] = useState({ title: "", overview: "" });


    useEffect(() => {
        setMovieList(JSON.parse(localStorage.getItem("my-list")));

        movieService.getGenres().then((data) => {
            setGenres(data.genres);
        })
    }, []);

    return (
        <div>
            <h1>Ma liste</h1>

            <Modal title={modalMovie.title} isActive={showModal} closeFunction={() => setShowModal(!showModal)} type="information">
                <p>{modalMovie.overview}</p>
            </Modal>
            {movieList && movieList.map((currrentMovie) => (
                <MovieCard movie={currrentMovie} key={currrentMovie.id} genres={genres} pshowModal={showModal} psetShowModal={setShowModal} psetModalMovie={setModalMovie} />
            ))}
        </div>
    );
}

export default Index;
