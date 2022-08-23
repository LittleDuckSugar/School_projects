import { useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image"
import movieService from "../services/movie.service";

const MovieCard = (props) => {

    // movieVideos saves datas about videos linked to a specific movie
    const [movieVideos, setMovieVideos] = useState({ results: [] });

    // loadVideo is a boolean: When true then the front will load a video, if false then nothing will be loaded
    const [loadVideo, setLoadVideo] = useState(false);

    // overLoad is a boolean: When true then the front will load over the picture, if false then nothing will be loaded
    const [overLoad, setOverLoad] = useState(false);

    // matchedGenre saves genre from current movie
    const [matchedGenre, setMatchedGenre] = useState();

    // showModal show the full movie informations available
    // const [showModal, setShowModal] = useState(false);

    // Saves videos informations about the movie
    useEffect(() => {
        movieService.getVideoFromMovieById(props.movie.id).then((data) => {
            setMovieVideos(data);
        });

        setMatchedGenre(props.genres && props.genres.filter((value) => {
            return props.movie.genre_ids.includes(value.id)
        }));
    }, []);

    // overLoader allow changes on the front by changing overLoad & loadVideo
    const overLoader = (state) => {
        setOverLoad(state);
        if (state) {
            if (movieVideos.results.length != 0) {
                setLoadVideo(true);
            } else {
                setLoadVideo(false);
            }
        } else {
            setLoadVideo(state);
        }
    }


    const toogleMovieList = (movie) => {
        const myListArray = [];

        //Si j'ai déjà un ou des films dans mon localstorage
        if (localStorage.getItem("my-list")) {

            const localStorageMovie = JSON.parse(localStorage.getItem("my-list"));
            localStorageMovie.forEach((movie) => {
                myListArray.push(movie);
            });

            const indexOfExistingMovie = myListArray.findIndex((el) => el.id === movie.id);

            if (indexOfExistingMovie !== -1) {
                const filteredMovies = myListArray.filter((item) => item.id != movie.id);
                localStorage.setItem("my-list", JSON.stringify(filteredMovies));
            }
            else {
                myListArray.push(movie);
                localStorage.setItem("my-list", JSON.stringify(myListArray));
            }

        }
        //Si localstorage vide
        else {
            myListArray.push(movie);
            localStorage.setItem("my-list", JSON.stringify(myListArray));
        }
    };

    return (
        <>
            <div className="movie__card" onMouseOver={() => overLoader(true)} onMouseLeave={() => overLoader(false)}>
                {overLoad ?

                    <div className="movie__data">

                        {loadVideo ?
                            <div className="movie__video">
                                <iframe width="560" height="315" src={`https://www.youtube.com/embed/${movieVideos.results[0].key}?controls=0&autoplay=1`} title={movieVideos.results[0].name} frameBorder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowFullScreen></iframe>
                            </div>
                            :
                            <div className="movie__img">
                                <img src={`https://image.tmdb.org/t/p/original${props.movie.backdrop_path}`} />
                            </div>
                        }

                        <div className="movie__about">
                            <h2>{props.movie.title}</h2>

                            <Link href={`/watch?movie=${props.movie.id}`}>
                                <Image src="/play.png" width="50px" height="50px" alt="Bouton play" />
                            </Link>

                            <Image src="/more.png" width="50" height="50" alt="More" onClick={() => {
                                props.psetModalMovie(props.movie);
                                props.psetShowModal(!props.pshowModal);
                            }} />

                            <button type="button" onClick={() => toogleMovieList(props.movie)}>Ajouter à ma liste</button>

                            <p id="movie__recommandation">Recommandé à {props.movie.vote_average * 10} %</p>

                            {props.movie.adult ? <h1>18+</h1> : <></>}

                            <ul>
                                {matchedGenre && matchedGenre.map((genre) => (
                                    <li id={genre.id}>{genre.name}</li>
                                ))}
                            </ul>
                        </div>

                    </div> :
                    <div className="movie__img">
                        <img src={`https://image.tmdb.org/t/p/original${props.movie.backdrop_path}`} />
                    </div>
                }
            </div>
        </>

    );
}

export default MovieCard;
