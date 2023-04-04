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
        try {
        //   const url = new URL(request.url);
        //   const searchParams = url.searchParams;
        //   const id = searchParams.get('id');
            
          const [incomeTypeResponse, spendingTypeResponse] = await Promise.all([
            fetch(`/income/type`, {
              headers: {
                Accept: "application/json",    
            },
              method: "GET",
            }),
            fetch(`/spending/type`, {
              headers: {
                Accept: "application/json",
              },
              method: "GET",
            })
          ]);
            
          if (!incomeTypeResponse.ok) {
            const error = new Error(`Could not fetch the post. Status: ${incomeTypeResponse.statusText}`);
            error.status = incomeTypeResponse.status;
            throw error;
          }
          
          if (!spendingTypeResponse.ok) {
            const error = new Error(`Could not fetch the comments. Status: ${spendingTypeResponse.statusText}`);
            error.status = spendingTypeResponse.status;
            throw error;
          }
          
          const incomeTypes = await incomeTypeResponse.json();
          const spendingTypes = await spendingTypeResponse.json();
          return { incomeTypes, spendingTypes };
        } catch (error) {
          console.log(error);
          throw error;
        }
})  