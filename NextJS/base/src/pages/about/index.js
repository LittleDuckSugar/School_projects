import TitlePage from '../../components/TitlePage';

const Index = () => {
    return (
        <div>
            <TitlePage title="About" />
            <div className="square__grid">
                <div className="square__container">
                    <div className="square"></div>
                    <div className="square"></div>
                    <div className="square big"></div>
                    <div className="square"></div>
                    <div className="square"></div>
                    <div className="square"></div>
                    <div className="square big"></div>
                    <div className="square"></div>
                    <div className="square"></div>
                </div>
            </div>
        </div>
    );
}

export default Index;
