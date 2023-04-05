import { redirect } from "react-router";


export const homeLoader = () => {
    
    const token = localStorage.getItem('token');

    if (token) {
        return null
    } else {
        return (redirect("/signin"))
    }
}


export const typesData =( async () => {
  const token = localStorage.getItem('token');
      
//   const currentDate = new Date();
// const firstDayOfMonth = new Date(currentDate.getFullYear(), currentDate.getMonth(), 1);

// const formatDate = (date) => {
//   const year = date.getFullYear();
//   const month = String(date.getMonth() + 1).padStart(2, '0');
//   const day = String(date.getDate()).padStart(2, '0');
  
//   return `${year}-${month}-${day}`;
// };

// const firstDayOfMonthFormatted = formatDate(firstDayOfMonth);
// const currentDateFormatted = formatDate(currentDate);
  
  
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
})  