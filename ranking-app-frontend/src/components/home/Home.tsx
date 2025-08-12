import React, { useState } from 'react';
import SearchBar from './components/searchBar/SearchBar';
import PlaceList from './components/placeList/PlaceList';

const Home = () => {
  const [searchQuery, setSearchQuery] = useState('');

  const handleSearch = (query: string) => {
    setSearchQuery(query);
  };

  return (
    <div>
      <SearchBar onSearch={handleSearch} />
      <PlaceList searchQuery={searchQuery} />
    </div>
  );
};

export default Home;