import {
  createBrowserRouter, 
  createRoutesFromElements,
  redirect,
  Route, 
  RouterProvider,
} from 'react-router-dom'  
// pages
import SignIn from './routes/Sign-In/SignIn'
import SignUp from './routes/Sign-Up/SignUp';
import HomePage from './routes/Home-Page/HomePage'
// layouts
import RootLayout from './layouts/RootLayout/RootLayout';
// errorElement
import { ErrorPage } from './routes/Error-Page/ErrorPage';
import { PageNotFound } from './routes/Page-Not-Found/PageNotFound';
import { budgetData, debtData, homeLoader } from './api/loaders';
import BudgetPage from './routes/Budget-Page/BudgetPage';
import PlannerPage from './routes/Planner-Page/PlannerPage';
import DebtPage from './routes/Debt-Page/DebtPage';
import SettingsPage from './routes/Settings-Page/SettingsPage';
import AboutUsPage from './routes/About-Us-Page/AboutUsPage';
import ContactsPage from './routes/Contacts-Page/ContactsPage';

const router = createBrowserRouter(
  createRoutesFromElements(
<Route>
  <Route element={<RootLayout />} id="root" loader={homeLoader} errorElement={<ErrorPage/>}>
      <Route path="/" >
          <Route index loader={()=>redirect("/statistics")}/>
          <Route path="statistics" element={<HomePage/>}/>
          <Route path="budget" loader={budgetData} element={<BudgetPage/>}/>
          <Route path="planner" element={<PlannerPage/>}/>
          <Route path="debt" loader={debtData} element={<DebtPage/>}/>
          <Route path="settings" element={<SettingsPage/>}/>
          <Route path="aboutus" element={<AboutUsPage/>}/>
          <Route path="contacts" element={<ContactsPage/>}/>
      </Route>
      <Route path="*" element={<PageNotFound/>} />
  </Route>
      <Route path="signin" element={<SignIn/>} />
      <Route path="signup" element={<SignUp/>}  />
</Route>
  )
)

const App = () => {
  return (
    <RouterProvider router={router} />
  );
}

export default App;