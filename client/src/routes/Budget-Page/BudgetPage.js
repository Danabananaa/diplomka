import { Grid, Box, FormControl, InputLabel, Stack, FilledInput, InputAdornment, TextField, MenuItem, Typography, Button } from "@mui/material";
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import { useLoaderData } from "react-router";
import { sendIncome } from "../../api/Budget/SendIncome";
import { useState } from "react";
import { useNavigate } from "react-router";
import { sendSpending } from "../../api/Budget/SendSpending";
import BudgetTable from "../../components/Budget-Table/BudgetTable";

// Sorting data from the server and returning a single array
const mergeAndSortByDate = (incomeArray, spendingArray) => {
  const incomeData = incomeArray.length > 0 ? incomeArray.map((item) => ({
    ID: 'income_'+item.ID,
    date: item.date,
    description: item.description,
    amount: item.amount,
    type_id: item.income_type_id,
    type: "Income",
  })) : [];

  const spendingData = spendingArray.length > 0 ? spendingArray.map((item) => ({
    ID: 'spending_' + item.ID,
    date: item.date,
    description: item.description,
    amount: item.amount,
    type_id: item.spending_type_id,
    type: "Spending",
  })) : [];

  const mergedArray = [...incomeData, ...spendingData];

  return mergedArray.sort((a, b) => new Date(b.date) - new Date(a.date));
};

const BudgetPage = () => {
    const navigate = useNavigate();
    const query = new URLSearchParams(location.search);
    const { Types, Stats } = useLoaderData(); //Spending and Income Types from the Server
    // console.log(Types.Type_spending); //Array with spending types
    // console.log(Types.Type_income); // Array with income types
    // console.log(Stats.Spending); // Spending data
    // console.log(Stats.Income);  // Income data 
    
    const mergedData = mergeAndSortByDate(Stats.Income, Stats.Spending);
    // console.log(mergedData);
    const [income, setIncome] = useState();
    const [spending, setSpending] = useState();
    const [incomeType, setIncomeType] = useState(Types && Types.Type_income && Types.Type_income[0].ID ? Types.Type_income[0].ID: 1);
    const [spendingType, setSpendingType] = useState(Types && Types.Type_spending && Types.Type_spending[0].ID ? Types.Type_spending[0].ID: 1);
    const [incomeDescription, setIncomeDescription] = useState();
    const [spendingDescription, setSpendingDescription] = useState();

  //FILTERS
    const handleButtonClick = (e, value) => {
      // e.preventDefault();
      if (value) {
        // Update the URL with the new page number
        const newQuery = new URLSearchParams(query);
        newQuery.set('filter', value);
        navigate({ search: newQuery.toString() }); // navigate to update state values with useEffect
      } 
    };
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
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'space-evenly'
          }}
        >
          {/* BUDGET DATA BOX */}
          
          <Box
          sx={{
            height: '5%',
            width: '95%',
            // backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
            // boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
            borderRadius: '16px',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            // border: '1px solid rgba(0, 0, 0, 0.25)',
            position: 'relative'
          }}
        >
        <Stack direction="row" spacing={5}>
      {['Year', 'Month', 'Week', 'Day'].map((period) => (
        <Button
          key={period}
          variant="contained"
          color="secondary"
          onClick={(e) => handleButtonClick(e, period.toLowerCase())}
          sx={{ border: '1px solid black', fontWeight: '600', color: '#0C1017', width: '100px' }}
        >
          {period}
        </Button>
      ))}
    </Stack>

      </Box>
          
          <Box
            sx={{
              height: '90%',
              width: '95%',
              // boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
              // borderRadius: '16px',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              // border: '1px solid rgba(0, 0, 0, 0.25)',
              position: 'relative'
            }}
          >
              {/* TABLE */}
            <BudgetTable mergedData={mergedData}/>


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
                        startAdornment={<InputAdornment position="start">₸</InputAdornment>}
                    />
                </FormControl>
              {/* TOP TYPE */}
                <TextField
                    id="outlined-select-currency"
                    variant="filled"
                    select
                    label="Select"
                    onChange={(e)=>setIncomeType(e.target.value)}
                    value={incomeType}
                    
                    sx={{width: '80%'}}
                    >
                    {Types.Type_income && Types.Type_income.map((option) => (
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
                        startAdornment={<InputAdornment position="start">₸</InputAdornment>}
                    />
                </FormControl>
              {/* BOTTOM TYPE */}
                <TextField
                    id="outlined-select-spending"
                    variant="filled"
                    select
                    onChange={(e)=>setSpendingType(e.target.value)}
                    value={spendingType}
                    // defaultValue={Types && Types.Type_spending && Types.Type_spending[0].ID ? Types.Type_spending[0].ID: 1}
                    label="Select"
                    sx={{width: '80%'}}
                >
                    {Types.Type_spending && (Types.Type_spending.map((option) => (
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