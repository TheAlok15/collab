import Header from '../components/Header.jsx';
import Sidebar from '../components/Sidebar.jsx';
import MainContent from '../components/MainContent.jsx';

const Home = () => {
  return (
    <div className="min-h-screen flex flex-col font-sans bg-bg-main selection:bg-tint-amber/30 selection:text-accent-amber">
      <Header />

      <div className="flex flex-1 pt-[52px]">
        
        <Sidebar />

        <div className="flex-1 ml-[260px] h-[calc(100vh-52px)]">
          <MainContent />
        </div>
      </div>
    </div>
  );
};

export default Home;