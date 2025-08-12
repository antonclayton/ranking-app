import { Link } from 'react-router-dom';
import styles from './Navbar.module.css';
import { useState } from 'react';

const Navbar = () => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <nav className={styles.navbar}>
      <div className={styles.navContainer}>
        <Link to="/" className={styles.logo} onClick={() => setIsOpen(false)}>
          RankIt
        </Link>
        
        <button 
          className={`${styles.hamburger} ${isOpen ? styles.active : ''}`}
          onClick={() => setIsOpen(!isOpen)}
          aria-label="Toggle menu"
        >
          <span className={styles.bar}></span>
          <span className={styles.bar}></span>
          <span className={styles.bar}></span>
        </button>

        <div 
          className={`${styles.navLinks} ${isOpen ? styles.active : ''}`}
          onClick={() => setIsOpen(false)}
        >
          <Link to="/" className={styles.navLink}>Home</Link>
          <Link to="/rankings" className={styles.navLink}>Rankings</Link>
          {/* <Link to="/about" className={styles.navLink}>About</Link> */}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;