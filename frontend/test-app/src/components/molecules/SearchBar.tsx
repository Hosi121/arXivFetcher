import React from 'react';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';

const SearchBar = () => {
  return (
    <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
      <TextField
        label="URL"
        variant="outlined"
        fullWidth
      />
      <Button
        variant="contained"
        color="primary"
      >
        Search
      </Button>
    </Box>
  );
};

export default SearchBar;

