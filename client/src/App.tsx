import React from 'react';
import { Grid } from '@material-ui/core';
import Contacts from './components/contacts/Contacts';

function App() {
  return (
    <Grid container direction="row" justify="flex-start" alignItems="center">
      <Grid item xs={3}>
        <Contacts />
      </Grid>
    </Grid>
  );
}

export default App;
