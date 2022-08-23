import React, { Fragment, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import PropTypes from 'prop-types';
// import axios from 'axios';

import {
    Box,
    Drawer,
    Hidden,
    List,
    Typography,
    makeStyles
} from '@material-ui/core';

import LocalMallIcon from '@mui/icons-material/LocalMall';
import PersonIcon from '@mui/icons-material/Person';
import LogoutIcon from '@mui/icons-material/Logout';

import NavItem from './NavItem';

const items = [
    {
        href: '/admin/painting',
        icon: LocalMallIcon,
        title: 'Product list'
    },
    {

        href: '/admin/user',
        icon: PersonIcon,
        title: 'User list',
    }
];

const useStyles = makeStyles(() => ({
    mobileDrawer: {
        backgroundColor: 'var(--blue-color)',
        width: 256
    },
    desktopDrawer: {
        backgroundColor: 'var(--blue-color)',
        width: 256,
        top: 64,
        height: 'calc(100% - 64px)'
    },
    name: {
        marginTop: '15px',
        color: 'white',
        textAlign: 'center'
    },
}));

const Navbar = ({ onMobileClose, openMobile }) => {
    const classes = useStyles();
    const location = useLocation();
    // const history = useHistory();

    // const handleLogout = (e) => {
    //     e.preventDefault();
    //     axios
    //       .post('')
    //       .then((result) => {
    //         console.log('RES DATA :', result.data);
    //         history.push('/login');
    //       })
    //       .catch(err => {
    //         console.error(err)
    //       });
    //   };

    useEffect(() => {
        if (openMobile && onMobileClose === false) {
            onMobileClose();
        }
    }, [location.pathname, onMobileClose, openMobile]);

    const content = (
        <Box height="100%" display="flex" flexDirection="column" >
            <Box alignItems="center" display="flex" flexDirection="column" p={2} >
                <Typography className={classes.name} variant="h6">
                    Menu
                </Typography>
            </Box>
            <Box p={2}>
                <List>
                    {items.map((item) => (
                        <NavItem
                            href={item.href}
                            key={item.title}
                            title={item.title}
                            icon={item.icon}
                        />
                    ))}
                </List>
            </Box>
            <Box flexGrow={1} />
            <Box p={2} m={2} >
                <List>
                    <NavItem
                        href='/authentication'
                        key='Logout'
                        title='Logout'
                        icon={LogoutIcon}
                    // onClick={handleLogout}
                    />
                </List>
            </Box>
        </Box>
    );

    return (
        <Fragment>
            <Hidden mdUp>
                <Drawer
                    anchor="left"
                    classes={{ paper: classes.mobileDrawer }}
                    onClose={onMobileClose}
                    open={openMobile}
                    variant="temporary"
                >
                    {content}
                </Drawer>
            </Hidden>
            <Hidden smDown>
                <Drawer
                    anchor="left"
                    classes={{ paper: classes.desktopDrawer }}
                    open
                    variant="persistent"
                >
                    {content}
                </Drawer>
            </Hidden>
        </Fragment>
    );
};

Navbar.propTypes = {
    onMobileClose: PropTypes.func,
    openMobile: PropTypes.bool
};

Navbar.defaultProps = {
    onMobileClose: () => { },
    openMobile: false
};

export default Navbar;
