import React, { useState } from 'react';
import SearchBar from './components/searchBar/SearchBar';
import PlaceList from './components/placeList/PlaceList';
import CreatePlaceModal from './components/PlaceModals/CreatePlaceModal';
import { Button } from '@mui/material';
import styles from './Home.module.css';
import AddIcon from '@mui/icons-material/Add';

const Home = () => {
  const [searchQuery, setSearchQuery] = useState('');
  const [showCreate, setShowCreate] = useState(false);
  const [refreshToken, setRefreshToken] = useState(0);

  const handleSearch = (query: string) => {
    setSearchQuery(query);
  };

  return (
    <div>
      <div className={styles.header}>
        <div className={styles.searchBarCol}>
          <SearchBar onSearch={handleSearch} />
        </div>
        <div className={styles.buttonCol}>
          <Button
            variant="contained"
            color="primary"
            onClick={() => setShowCreate(true)}
            className={styles.createButton}
            sx={{
              fontWeight: 700,
            }}
            endIcon={<AddIcon fontSize="large" />}
          >
            Add Place
          </Button>
        </div>
      </div>
      <PlaceList
        searchQuery={searchQuery}
        refreshToken={refreshToken}
        onEdited={() => setRefreshToken((t) => t + 1)}
      />

      <CreatePlaceModal
        open={showCreate}
        onClose={() => setShowCreate(false)}
        onCreated={() => setRefreshToken((t) => t + 1)}
      />
    </div>
  );
};

export default Home;