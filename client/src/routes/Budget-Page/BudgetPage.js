import { Grid, Box, FormControl, InputLabel, FilledInput, InputAdornment, TextField, MenuItem, Typography, Button } from "@mui/material";
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import { useLoaderData } from "react-router";
import { sendIncome } from "../../api/Budget/SendIncome";
import { useState } from "react";
import { useNavigate } from "react-router";
import { sendSpending } from "../../api/Budget/SendSpending";

const BudgetPage = () => {
    const navigate = useNavigate();
    const { incomeTypes, spendingTypes } = useLoaderData(); //Spending and Income Types from the Server
    const [income, setIncome] = useState();
    const [spending, setSpending] = useState();
    const [incomeType, setIncomeType] = useState();
    const [spendingType, setSpendingType] = useState();
    const [incomeDescription, setIncomeDescription] = useState();
    const [spendingDescription, setSpendingDescription] = useState();
    return (
        <Grid 
            container 
            alignItems="center"
            sx={{ 
            height: '100%'
            }}
        >
      {/* BUDGET MAIN CONTAINER */}
        <Grid item xs={7} 
          sx={{
            height: '100%',
            // border:'1px solid black',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
          }}
        >
          {/* BUDGET DATA BOX */}
          <Box
            sx={{
              height: '95%',
              width: '95%',
              backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
              boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
              borderRadius: '16px',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              border: '1px solid rgba(0, 0, 0, 0.25)',
              position: 'relative'
            }}
          >
            {/* BUDGET DATA */}

          </Box>
        </Grid>
        {/* RIGHT CONTAINER */}
        <Grid item xs={5} 
          sx={{ 
            height: '100%',
            // border: '1px solid black',
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'space-evenly',
            alignItems: 'center'
          }}
        >
          {/* TOP BOX */}
          <Box
              sx={{
                height: '45%',
                width: '80%',
                backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
                boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
                borderRadius: '16px',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'space-evenly',
                border: '1px solid rgba(0, 0, 0, 0.25)'
              }}
            >
                {/* TOP HEADER */}
                <Typography variant="h4" sx={{fontFamily: 'DejaVuSans'}}>
                    Кірісті қосу
                </Typography>
                {/* TOP AMOUNT */}
                <FormControl sx={{ width: '80%' }} variant="filled">
                    <InputLabel htmlFor="filled-adornment-amount">Amount</InputLabel>
                    <FilledInput
                        id="filled-adornment-amount"
                        onChange={(e)=>setIncome(e.target.value)}
                        startAdornment={<InputAdornment position="start">$</InputAdornment>}
                    />
                </FormControl>
              {/* TOP TYPE */}
                <TextField
                    id="outlined-select-currency"
                    variant="filled"
                    select
                    label="Select"
                    onChange={(e)=>setIncomeType(e.target.value)}
                    defaultValue={incomeTypes[0].IncomeType ? incomeTypes[0].IncomeType: ''}
                    sx={{width: '80%'}}
                    >
                    {incomeTypes && incomeTypes.map((option) => (
                        <MenuItem key={option.ID} value={option.ID}>
                            {option.IncomeType}
                        </MenuItem>
                    ))}
                </TextField>
                {/* Description */}
                <TextField 
                    id="filled-basic" 
                    label="Description" variant="filled" 
                    onChange={(e)=> setIncomeDescription(e.target.value)}
                    sx={{
                        width: '80%'
                    }}
                />
            {/* TOP BUTTON */}
                <AddCircleOutlineIcon 
                    onClick={(e) => sendIncome(e, incomeType, income, incomeDescription, navigate)}
                    sx={{
                        fontSize:'50px',
                        '&:hover':{
                            cursor: 'pointer',
                            color: 'white'
                        }
                    }}
                />

          </Box>
          {/* BOTTOM BOX */}
          <Box
              sx={{
                height: '45%',
                width: '80%',
                backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
                boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
                borderRadius: '16px',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'space-evenly',
                border: '1px solid rgba(0, 0, 0, 0.25)'
              }}
            >
              {/* BOTTOM HEADER */}
                <Typography variant="h4">
                    Шығысты қосу
                </Typography>
                {/* BOTTOM AMOUNT */}
                <FormControl sx={{ width: '80%' }} variant="filled">
                    <InputLabel htmlFor="filled-adornment-amount">Amount</InputLabel>
                    <FilledInput
                        id="filled-adornment-amount"

                        onChange={(e)=>setSpending(e.target.value)}
                        startAdornment={<InputAdornment position="start">$</InputAdornment>}
                    />
                </FormControl>
              {/* BOTTOM TYPE */}
                <TextField
                    id="outlined-select-spending"
                    variant="filled"
                    select
                    onChange={(e)=>setSpendingType(e.target.value)}
                    label="Select"
                    defaultValue={ spendingTypes[0].SpendingType? spendingTypes[0].SpendingType: ''}
                    sx={{width: '80%'}}
                >
                    {spendingTypes && (spendingTypes.map((option) => (
                        // console.log(option.SpendingType)
                        <MenuItem key={option.ID} value={option.ID}>
                            {option.SpendingType}
                        </MenuItem>
                        ))
                    )}
                </TextField>
                {/* Description */}
                <TextField 
                    id="filled-basic" 
                    label="Description" 
                    variant="filled" 
                    onChange={(e)=> setSpendingDescription(e.target.value)}
                    sx={{
                        width: '80%'
                    }}
                />
            {/* BOTTOM BUTTON */}
                <AddCircleOutlineIcon 
                    onClick={(e)=>sendSpending(e, spendingType, spending, spendingDescription, navigate)}
                    sx={{
                        fontSize:'50px',
                        '&:hover':{
                            cursor: 'pointer',
                            color: 'white'
                        }
                    }}
                />
          </Box>
        </Grid>
      </Grid>
      
    )
}
export default BudgetPage;