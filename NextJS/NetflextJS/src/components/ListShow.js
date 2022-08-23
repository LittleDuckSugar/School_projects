import React from 'react';
import MovieCard from './MovieCard';

const ListShow = (props) => {
    return (
        <div>
            <h1>{props.listName}</h1>
            <div className="movies__grid">
                {props.movies && props.movies.map((currrentMovie) => (
                    <MovieCard movie={currrentMovie} key={currrentMovie.id} genres={props.genres} pshowModal={props.showModal} psetShowModal={props.setShowModal} psetModalMovie={props.setModalMovie} />
                ))}
            </div>
        </div>
    );
}

export default ListShow;
