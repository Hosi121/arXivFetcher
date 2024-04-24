import React from 'react';
import TextField from '@mui/material/TextField';

const URLInput = ({ ...props }) => {
  return (
    <TextField
      label="URL"
      variant="outlined"
      fullWidth
      {...props} // 他のプロパティがあれば展開する
    />
  );
};

export default URLInput;
