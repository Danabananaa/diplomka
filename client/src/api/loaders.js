import { redirect } from "react-router";

const spendingColors = ['DC2F02', 'E85D04', 'F48C06', 'FAA307', 'D00000'];
const incomeColors = ['9EF01A', '70E000', 'CCFF33', '38B000'];


function convertToPieChartData(financialData) { // CONVERTS loader data on page Statistics to the format that is acceptable with the Donut Chart
  const pieChartData = [
    ...financialData.spendings.map((s, index) => ({
      id: `spending-${s.typeId}`,
      label: `Spending Type ${s.typeId}`,
      value: s.total,
      color: `#${spendingColors[index % spendingColors.length]}`,
    })),
    ...financialData.incomes.map((i, index) => ({
      id: `income-${i.typeId}`,
      label: `Income Type ${i.typeId}`,
      value: i.total,
      color: `#${incomeColors[index % incomeColors.length]}`,
    })),
  ];

  return pieChartData;
} 


export const homeLoader = () => {
    
    const token = localStorage.getItem('token');

    if (token) {
        return null
    } else {
        return (redirect("/signin"))
    }
}

export const statisticsData =( async ({request}) => {
  const token = localStorage.getItem('token');
  
  if (token) { //If token is valid do the fetch  
  const url = new URL(request.url);
  const currentQuery = url.search;
  
  try {
    const [budgetTypeResponse,statResponse] = await Promise.all([
      fetch(`/budget/type`, {
        headers: {
          Accept: "application/json",
          Authorization: `${token}`,
        },
        method: "GET",
        
      }),
      fetch(`/statistics${currentQuery}`, {
        headers: {
          Accept: "application/json",
          Authorization: `${token}`,
        },
        method: "GET",
      }),
    ]);
    if (!budgetTypeResponse.ok) {
      const error = new Error(`Could not fetch posts. Status: ${budgetTypeResponse.statusText}`);
      error.status = budgetTypeResponse.status;
      throw error;
    }
    if (!statResponse.ok) {
      const error = new Error(`Could not fetch posts. Status: ${statResponse.statusText}`);
      error.status = statResponse.status;
      throw error;
    }
    const types = await budgetTypeResponse.json(); //Data from budgetTypeResponse
    
    const data = await statResponse.json(); // Data from statResponse
    const financialData = null
   // Create a mapping from typeId to IncomeType and SpendingType
   if (types.Type_income && types.Type_spending){

     const incomeTypeMapping = types.Type_income.reduce((acc, type) => {
      acc[type.ID] = type.IncomeType;
      return acc;
    }, {});
  
    const spendingTypeMapping = types.Type_spending.reduce((acc, type) => {
      acc[type.ID] = type.SpendingType;
      return acc;
    }, {});
    if (data.value_spending.spendings && data.value_income.incomes && data.value_income.total_amount && data.value_spending.total_amount){
      financialData = {
        spendings: data.value_spending.spendings.map(s => ({
          typeId: spendingTypeMapping[s.type_id], // Use mapping here
          percentage: s.percentage,
          total: s.total,
        })),
        totalSpending: data.value_spending.total_amount,
        incomes: data.value_income.incomes.map(i => ({
          typeId: incomeTypeMapping[i.type_id], // Use mapping here
          percentage: i.percentage,
          total: i.total,
        })),
        totalIncome: data.value_income.total_amount,
      };
    }
  } else{
    throw new Error(`Couldn't fetch types data`)
  }
  
  let pieChartData = null;
  if (financialData){
    pieChartData = convertToPieChartData(financialData);
  }
    return { pieChartData };
  } catch (error) {
    throw error;
  }
} else {
    return (redirect("/signin"))
}
  
  
  

  // try{
  //   const response = await fetch(`/statistics`, {
  //     headers: {
  //       Accept: 'application/json',
  //       Authorization: `${token}`,
  //       },
  //     method: "GET",
  //   })
  //   if (!response.ok){
      
  //       const error = new Error(response.statusText);
  //       error.status = response.status;
  //       throw error;
  //   } else if (response.ok){
  //     const data = await response.json();
  //     console.log(data);
  //     return {data};
  //   }
  // } catch(error){
  //   console.log(error);
  //   throw error;
  // }
})  //OK

export const budgetData =( async () => {
  const token = localStorage.getItem('token');
  
  if (token) { // Do fetch if token is present

    try {
      //   const url = new URL(request.url);
      //   const searchParams = url.searchParams;
      //   const id = searchParams.get('id');
          
        const [budgetTypeResponse, budgetStatsResponse] = await Promise.all([
          fetch(`/budget/type`, {
            headers: {
              Accept: "application/json",
              Authorization: `${token}`,
            },
            method: "GET",
            
          }),
          fetch(`/budget/stats`, {
            headers: {
              Accept: "application/json",
              Authorization: `${token}`,
            },
            method: "GET",
            
          })
        ]);
          
        if (!budgetTypeResponse.ok) {
          // if (budgetTypeResponse.status===401){
          //   return redirect('/signin')
          // }
          const error = new Error(`Could not fetch the post. Status: ${budgetTypeResponse.statusText}`);
          error.status = budgetTypeResponse.status;
          throw error;
        }
        
        if (!budgetStatsResponse.ok) {
          // if (budgetStatsResponse.status===401){
          //   return redirect('/signin')

          // }
          const error = new Error(`Could not fetch the comments. Status: ${budgetStatsResponse.statusText}`);
          error.status = budgetStatsResponse.status;
          throw error;
        }
        
        const Types = await budgetTypeResponse.json();
        const Stats = await budgetStatsResponse.json();
        return { Types, Stats };
      } catch (error) {
        console.log(error);
        throw error;
      }

  } else {
      return (redirect("/signin"))
  }  
})  

export const debtData =( async () => {
  const token = localStorage.getItem('token');
  try {
        //   const url = new URL(request.url);
        //   const searchParams = url.searchParams;
        //   const id = searchParams.get('id');
            
          const [debtResponse, debtTypesResponse] = await Promise.all([
            fetch(`/debt`, {
              headers: {
                Accept: "application/json",
                Authorization: `${token}`,
              },
              method: "GET",
              
            }),
            fetch(`/debt/type`, {
              headers: {
                Accept: "application/json",
                Authorization: `${token}`,
              },
              method: "GET",
              
            }),
          ]);
            
          if (!debtResponse.ok) {
            // if (budgetTypeResponse.status===401){
            //   return redirect('/signin')
            // }
            const error = new Error(`Could not fetch the post. Status: ${budgetTypeResponse.statusText}`);
            error.status = budgetTypeResponse.status;
            throw error;
          }
          if (!debtTypesResponse.ok) {
            // if (budgetTypeResponse.status===401){
            //   return redirect('/signin')
            // }
            const error = new Error(`Could not fetch the post. Status: ${budgetTypeResponse.statusText}`);
            error.status = budgetTypeResponse.status;
            throw error;
          }
          
          const debtData = await debtResponse.json();
          const debtTypesData = await debtTypesResponse.json();
          // console.log(debtTypesData);
          console.log(debtData);
          return { debtData, debtTypesData };
        } catch (error) {
          console.log(error);
          throw error;
        }
})