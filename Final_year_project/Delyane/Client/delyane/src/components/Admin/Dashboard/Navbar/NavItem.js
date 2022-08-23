import React from 'react';
import PropTypes from 'prop-types';
import { NavLink } from 'react-router-dom';

import clsx from 'clsx';

import {
    Button,
    ListItem,
    makeStyles
} from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
    item: {
        display: 'flex',
        paddingTop: 0,
        paddingBottom: 0
    },
    button: {
        color: theme.palette.text.secondary,
        fontWeight: theme.typography.fontWeightMedium,
        justifyContent: 'flex-start',
        letterSpacing: 0,
        padding: '10px 8px',
        textTransform: 'none',
        width: '100%'
    },
    icon: {
        color: 'var(--white-color)',
        marginRight: theme.spacing(1)
    },
    title: {
        color: 'var(--white-color)',
        marginRight: 'auto'
    },
    active: {
        color: 'var(--white-color)',
        background: 'linear-gradient(270deg, #F0F6FF 0%, #71A5EB 0.01%, #113D78 96.34%)',
        '& $title': {
            fontWeight: theme.typography.fontWeightBold
        },
        '& $icon': {
            transform: 'scale(1.3)',
            color: 'var(--white-color)'
        }
    }
}));

const NavItem = ({ className, href, icon: Icon, title, ...rest }) => {
    const classes = useStyles();

    return (
        <ListItem className={clsx(classes.item, className)} disableGutters {...rest}>
            <Button activeClassName={classes.active} className={classes.button} component={NavLink} to={href} >
                {Icon && (<Icon className={classes.icon} size="20" />)}
                <span className={classes.title}>
                    {title}
                </span>
            </Button>
        </ListItem>
    );
};

NavItem.propTypes = {
    className: PropTypes.string,
    href: PropTypes.string,
    icon: PropTypes.elementType,
    title: PropTypes.string
};

export default NavItem;