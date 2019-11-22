import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import {  NavLink } from "react-router-dom";
import { ListItemIcon, ListItemText } from '@material-ui/core';
import MailIcon from '@material-ui/icons/Mail';

const drawerWidth = 260;

const useStyles = makeStyles(theme => ({
  root: {
    display: 'flex',
  },
  appBar: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  toolbar: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.default,
    padding: theme.spacing(3),
  },
  listItem: {
    display: 'inline-flex', 
    textDecoration: 'none',
    color: 'rgba(0, 0, 0, 0.87)'
  }
}));


export const Sidebar = () => {
  const classes = useStyles();

  return (
      <Drawer
        className={classes.drawer}
        variant="permanent"
        classes={{
          paper: classes.drawerPaper,
        }}
        anchor="left"
      >
        <div className={classes.toolbar} style={{textAlign: 'center', marginTop: 20, minHeight: 20, color: 'darkcyan', fontSize: 20, fontWeight: 'bold'}}>
          Komodor.io
        </div>
        <List>
          <ListItem key="releases" component={NavLink} to="/releases" className={classes.listItem}>
              <ListItemIcon><MailIcon /></ListItemIcon>
              <ListItemText primary="Releases" />
          </ListItem>
          <ListItem button key="history"  component={NavLink} to="/history" className={classes.listItem}>
              <ListItemIcon><MailIcon /></ListItemIcon>
              <ListItemText primary="History" />
          </ListItem>
          <ListItem button key="pipelines" component={NavLink} to="/pipelines" className={classes.listItem}>
              <ListItemIcon><MailIcon /></ListItemIcon>
              <ListItemText primary="Pipelines" />
          </ListItem>
       </List>
      </Drawer>
  );
}
