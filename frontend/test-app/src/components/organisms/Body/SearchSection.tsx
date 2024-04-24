import React from 'react';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import Typography from '@mui/material/Typography';
import SearchBar from './SearchBar';

const SearchSection = () => {
  return (
    <div>
      <SearchBar />
      <Typography variant="h6" sx={{ mt: 2 }}>
        How to use?
      </Typography>
      <List>
        <ListItem>Enter the URL you wish to search.</ListItem>
        <ListItem>Press the "Search" button to proceed.</ListItem>
        <ListItem>View the search results below.</ListItem>
      </List>
    </div>
  );
};

export default SearchSection;

