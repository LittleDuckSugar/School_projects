import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import movieService from "../../services/movie.service";

const Index = () => {
    const [movieVideos, setMovieVideos] = useState({results: []});

    const router = useRouter();

    useEffect(() => {
        if (!router.isReady) return;
        console.log("The id of the movie asked is :", router.query.movie);
        movieService.getVideoFromMovieById(router.query.movie).then((data) => {
            setMovieVideos(data);
        }).catch((err) => console.log(err));
    }, [router.isReady]);

    return (
        <div className="movie__watcher">

            {movieVideos.results.length != 0 ?
                <div className="movie__video">
                    <p>Movies are not available here, but here is a video</p>
                    <iframe width="560" height="315" src={`https://www.youtube.com/embed/${movieVideos.results[0].key}?autoplay=1`} title={movieVideos.results[0].name} frameBorder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowFullScreen></iframe>
                </div>
                :
                <div className="movie__replacement">
                    <p>No videos for this movie, but here is a great song</p>
                    <iframe width="560" height="315" src="https://www.youtube.com/embed/dQw4w9WgXcQ?controls=0&autoplay=1" title="Rickrolled" frameBorder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowFullScreen></iframe>
                </div>
            }
        </div>
    );
}

export default Index;
