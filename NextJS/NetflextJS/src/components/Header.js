import Link from "next/link";

const Header = () => {
    return (
        <header className="header__main">
            <div className="header__logo">
                <img src="https://image.tmdb.org/t/p/original/wwemzKWzjKYJFfCeiB57q3r4Bcm.svg" alt="Netflix" />
            </div>
            <nav className="header__nav">
                <ul className="nav__list">
                    <li className="nav__item">
                        <Link href="/">
                            <a className="nav__link">Home</a>
                        </Link>
                    </li>
                    <li className="nav__item">
                        <Link href="/YourAccount">
                            <a className="nav__link">Compte</a>
                        </Link>
                    </li>
                    <li className="nav__item">
                        <Link href="/browse/my-list">
                            <a className="nav__link">Ma liste</a>
                        </Link>
                    </li>
                </ul>
            </nav>
        </header>
    );
}

export default Header;
