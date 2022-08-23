import React from 'react';
import { Route, Switch } from 'react-router-dom';

import './App.css';

import Admin from './components/Admin/Admin/Admin';
import Authentication from './components/Authentication/Authentication';
import CreateCustomer from './components/Admin/Dashboard/Dashboard/CreateCustomer';
import CreatePainting from './components/Admin/Dashboard/Dashboard/CreatePainting';
import Customer from './components/Admin/Dashboard/Dashboard/Customer';
import EditCustomer from './components/Admin/Dashboard/Dashboard/EditCustomer';
import Error from './components/Error/Error';
import Home from './components/Home/Home';
import Favorite from './components/Favorite/Favorite';
import Painting from './components/Admin/Dashboard/Dashboard/Painting';
import Product from './components/Product/[uuid]/Product';
import Products from './components/Product/Products';
import Register from './components/Register/Register';
import Profil from './components/Profil/Profil';
import Cart from './components/Cart/Cart';

import { library } from "@fortawesome/fontawesome-svg-core";
import { faBasketShopping, faUser, faHeart, faChevronRight, faChevronLeft, faHeartCirclePlus } from "@fortawesome/free-solid-svg-icons";
import { faCcPaypal, faCcVisa, faCcMastercard } from '@fortawesome/free-brands-svg-icons'
library.add(faBasketShopping, faUser, faHeart, faCcPaypal, faCcVisa, faCcMastercard, faChevronRight, faChevronLeft, faHeartCirclePlus);

const App = () => {
  return (
    <div className="App">
      <Switch >
        <Route exact path='/' component={Home} />
        <Route exact path='/admin' component={Admin} />
        <Route exact path='/admin/user' component={Customer} />
        <Route exact path='/admin/newcustomer' component={CreateCustomer} />
        <Route exact path='/admin/newpainting' component={CreatePainting} />
        <Route exact path='/admin/editcustomer/:uuid' component={EditCustomer} />
        <Route exact path='/admin/painting' component={Painting} />
        <Route exact path='/authentication' component={Authentication} />
        <Route exact path='/favorite' component={Favorite} />
        <Route exact path='/cart' component={Cart} />
        <Route exact path='/painting' component={Products} />
        <Route exact path='/painting/:uuid' component={Product} />
        <Route exact path='/register' component={Register} />
        <Route exact path='/profil' component={Profil} />
        <Route component={Error} />
      </Switch>
    </div>
  );
}

export default App;
