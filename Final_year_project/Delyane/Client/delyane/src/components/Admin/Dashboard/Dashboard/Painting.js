import React, { useState, useEffect } from 'react';
import axios from 'axios';
import clsx from 'clsx';

import Headers from '../Navbar/Headers';

import {
    Button,
    Card,
    Divider,
    Grid,
    InputAdornment,
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TablePagination,
    TableRow,
    TextField,
    Typography,
} from '@material-ui/core';

import SearchIcon from '@material-ui/icons/Search';

import { makeStyles } from '@material-ui/core/styles';

const columns = [
    { id: 'title', label: 'Title', minWidth: 100, align: 'center' },
    { id: 'description', label: 'Description', minWidth: 100, align: 'center' },
    { id: 'price', label: 'Price', minWidth: 100, align: 'center' },
    { id: 'view', label: '', minWidth: 100, align: 'center' },
];

const useStyles = makeStyles((theme) => ({
    root: {
        backgroundColor: 'var(--lightgray-color)',
        position: 'absolute',
        display: 'flex',
        height: '100%',
        overflow: 'hidden',
        width: '100%',
    },
    wrapper: {
        display: 'flex',
        flex: '1 1 auto',
        overflow: 'hidden',
        paddingTop: 64,
        [theme.breakpoints.up('md')]: {
            paddingLeft: 256,
        },
        [theme.breakpoints.up('lg')]: {
            paddingLeft: 256,
        },
    },
    contentContainer: {
        display: 'flex',
        flex: '1 1 auto',
        overflow: 'hidden',
    },
    content: {
        flex: '1 1 auto',
        height: '100%',
        overflow: 'auto',
    },
    pageContainer: {
        minHeight: '100%',
        marginTop: '0px',
    },
    secondRoot: {
        margin: '20px 20px',
        marginBottom: '50px',
        padding: '20px 20px',
    },
    mainTable: {
        margin: '20px 0 20px',
    },
    title: {
        color: 'var(--blue-color)',
        textAlign: 'center',
    },
    container: {
        maxHeight: 440,
    },
    space: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        margin: '20px 20px',
        [theme.breakpoints.down('xs')]: {
            display: 'block',
            textAlign: 'center',
        },
    },
    search: {
        verticalAlign: 'baseline',
    },
    searchIc: {
        color: 'var(--lightgray-color)',
    },
    button: {
        background: 'var(--blue-color)',
        color: 'var(--white-color)',
        margin: '15px',
        '&:hover': {
            backgroundColor: '#113e78ec',
        },
    },
    divider: {
        backgroundColor: 'var(--blue-color)',
        height: '5px',
        margin: '10px 0',
    },
    btn: {
        display: 'flex',
        justifyContent: 'center',
    },
    viewbtn: {
        color: 'var(--blue-color)',
        background: 'var(--white-color)',
        border: '2px solid',
        margin: '10px',
        minInlineSize: 'auto',
        padding: '5px',
    }
}));

const Painting = ({ className, staticContext, ...rest }) => {
    const classes = useStyles();
    const [page, setPage] = useState(0);
    const [products, setProducts] = useState([]);

    const [searchTerm, setSearchTerm] = useState('');
    const [rowsPerPage, setRowsPerPage] = useState(10);

    useEffect(() => {
        const getDatas = async () => {
            try {
                const result = await axios.get('http://90.22.250.124:8080/products')
                setProducts(result.data);
            } catch (err) {
                console.log(err)
            }
        };
        getDatas();
    }, []);

    const handleChangePage = (newPage) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(+event.target.value);
        setPage(0);
    };

    const tableHeader = columns.map((column) => (
        <TableCell key={column.id} align={column.align} style={{ minWidth: column.minWidth }}>
            {column.label}
        </TableCell>
    ));

    const mainContent = products
        .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
        .filter(
            (product) =>
                !searchTerm || product.title.toLowerCase().includes(searchTerm.toLowerCase())
        )
        .map((product) => {
            return (
                <TableRow hover role="checkbox" tabIndex={-1} key={product.uuid}>
                    <TableCell align="center">{product.title}</TableCell>
                    <TableCell align="center">{product.description}</TableCell>
                    <TableCell align="center">{product.price} €</TableCell>
                    <TableCell align="center">
                        <Button className={classes.button} href={`/painting/${product.uuid}`}>View</Button>
                        <Button className={classes.button} href={`/admin/editproduct/${product.uuid}`}>Edit</Button>
                    </TableCell>
                </TableRow>
            );
        });

    return (
        <div className={classes.root}>
            <div className={classes.wrapper}>
                <div className={classes.contentContainer}>
                    <div className={classes.content}>

                        <Headers />

                        <div className={clsx(classes.pageContainer, className)}>
                            <Card className={clsx(classes.secondRoot, className)} elevation={2} {...rest}>
                                <Typography className={classes.title} gutterBottom variant="h4" component="h1">
                                    Products list
                                </Typography>

                                <div className={classes.space}>
                                    <TextField className={classes.search} id="outlined-search" label="Search" type="search"
                                        onChange={(event) => { setSearchTerm(event.target.value) }}
                                        InputProps={{
                                            endAdornment: (
                                                <InputAdornment position="end">
                                                    <SearchIcon className={classes.searchIc} />
                                                </InputAdornment>
                                            ),
                                        }}
                                    />
                                    <Button className={classes.button} href='/admin/newpainting'>
                                        Create a new product
                                    </Button>
                                </div>

                                <Divider className={classes.divider} />

                                <Grid>
                                    <Paper className={classes.mainTable}>
                                        <TableContainer className={classes.container}>
                                            <Table stickyHeader aria-label="sticky table">
                                                <TableHead>
                                                    <TableRow>{tableHeader}</TableRow>
                                                </TableHead>
                                                <TableBody>{mainContent}</TableBody>
                                            </Table>
                                        </TableContainer>
                                        <TablePagination
                                            rowsPerPageOptions={[10, 25, 100]}
                                            component="div"
                                            count={mainContent.length}
                                            rowsPerPage={rowsPerPage}
                                            page={page}
                                            onPageChange={handleChangePage}
                                            onRowsPerPageChange={handleChangeRowsPerPage}
                                        />
                                    </Paper>
                                </Grid>
                            </Card>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Painting;
