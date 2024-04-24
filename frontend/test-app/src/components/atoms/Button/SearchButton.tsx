import React from 'react';
import Button from '@mui/material/Button';

const SearchButton = ({ ...props }) => {
  return (
    <Button
      variant="contained"
      color="primary"
      {...props} // 他のプロパティがあれば展開する
    >
      Search
    </Button>
  );
};

export default SearchButton;

