import { Outlet } from 'react-router-dom';
import Navbar from '../navbar/Navbar';
import styles from './Layout.module.css';

const Layout = () => {
  return (
    <div className={styles.appContainer}>
      <Navbar />
      <div className={styles.pageContent}>
        {/* The Outlet is where the nested page components will be rendered */}
        <Outlet />
      </div>
    </div>
  );
};

export default Layout;