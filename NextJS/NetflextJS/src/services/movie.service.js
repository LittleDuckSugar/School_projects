const apiURL = "https://api.themoviedb.org/";

const v3 = "3";
const v4 = "4";

// const apiKeyV3 = "99357a8604e4f930825d5ab16a4b3b72"; // Used in params
const apiKeyV4 = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI5OTM1N2E4NjA0ZTRmOTMwODI1ZDVhYjE2YTRiM2I3MiIsInN1YiI6IjYyMmYzY2MwMTI5NzBjMDA0NjlkMzQ1NiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.mtJqlbhe46H7DW2KbfXKIUQtCTGLCyXDnKI4LO70rXo"; // Bearer token

export default {
    getVideoFromMovieById(id) {
        return fetch(`${apiURL}${v3}/movie/${id}/videos`, {
            headers: {
                "Authorization": `Bearer ${apiKeyV4}`,
                "Content-Type": "application/json;charset=utf-8"
            }
        }).then(res => res.json());
    },

    getDiscoveryMovies() {
        return fetch(`${apiURL}${v4}/discover/movie?page=${Math.floor(Math.random() * 500) + 1}`, {
            headers: {
                "Authorization": `Bearer ${apiKeyV4}`,
                "Content-Type": "application/json;charset=utf-8"
            }
        }).then(res => res.json());
    },

    getPopularMovies() {
        return fetch(`${apiURL}${v3}/movie/popular`, {
            headers: {
                "Authorization": `Bearer ${apiKeyV4}`,
                "Content-Type": "application/json;charset=utf-8"
            }
        }).then(res => res.json());
    },

    getGenres() {
        return fetch(`${apiURL}${v3}/genre/movie/list`, {
            headers: {
                "Authorization": `Bearer ${apiKeyV4}`,
                "Content-Type": "application/json;charset=utf-8"
            }
        }).then(res => res.json());
    }
}

