import React from 'react';
import { Typography } from '@mui/material';
import {Box} from '@mui/material';
import DonutChart from '../../components/DonutChart/DonutChart';
const HomePage = () => {
  return (
    
      <Box>

        <Typography variant="h4" gutterBottom>
          Overall budget
        </Typography>
        
        <DonutChart/>

      </Box>
    
  );
};

export default HomePage;