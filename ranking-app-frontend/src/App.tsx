import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Layout from './components/layout/Layout';
import Home from './components/home/Home';
import Rankings from './components/rankings/Rankings';
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          {/* These are the pages that will be rendered within the <Outlet> */}
          <Route index element={<Home />} />
          <Route path="rankings" element={<Rankings />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;