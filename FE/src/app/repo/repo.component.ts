import { Component, OnInit } from '@angular/core';
import { Store, select } from '@ngrx/store';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { Router } from '@angular/router';
import { FormControl } from '@angular/forms';

import { Repo } from '../repo.model';
import RepoState from '../repo.state';
import { getTestBed } from '@angular/core/testing';
import * as RepoActions from '../repo.action';
import TestedRepoModel from '../repo.model';

@Component({
  selector: 'repo',
  templateUrl: './repo.component.html',
  styleUrls: ['./repo.component.scss'],
})
export class RepoComponent implements OnInit {
  inputRepo = new FormControl('');
  constructor(private store: Store<RepoState>, private route: Router) {}

  ngOnInit(): void {}

  getTested() {
    let userInputList = this.inputRepo.value.split('/');
    if (userInputList.length != 2) {
      alert('Invalid input, Please try another github repo.');
      return;
    }
    const repo: Repo = { owner: userInputList[0], repoName: userInputList[1] };
    this.store.dispatch(RepoActions.GetRepoAction({ payload: repo }));
    this.route.navigate(['/files']);
  }
}
