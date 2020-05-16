import { Component, OnInit } from '@angular/core';
import { Store, select } from '@ngrx/store';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { Router } from '@angular/router';

import { Repo, FileModel, ModifiedFileModel } from '../repo.model';
import RepoState from '../repo.state';
import { getTestBed } from '@angular/core/testing';
import * as RepoActions from '../repo.action';
import TestedRepoModel from '../repo.model';
import { THIS_EXPR } from '@angular/compiler/src/output/output_ast';

@Component({
  selector: 'files',
  templateUrl: './files.component.html',
  styleUrls: ['./files.component.scss'],
})
export class FilesComponent implements OnInit {
  repoSub$: Observable<RepoState>;
  repoObject: RepoState;
  repo: TestedRepoModel;
  modifiledFiles: ModifiedFileModel[];
  currentFile: ModifiedFileModel;
  displayedColumns: string[] = ['name'];

  test1 = 'fdhkadjkaf';
  test2 = '        dfdafd';

  constructor(
    private store: Store<{ repo: RepoState }>,
    private route: Router
  ) {
    this.repoSub$ = this.store.select('repo').pipe(
      tap((result) => {
        this.repoObject = result;
        this.repo = this.repoObject.tested;
        if (this.repo) {
          this.modifiledFiles = this.repo.goFiles.map((x) => {
            let splited = x.content.split('\n');
            let splitedContent = []

            for (let i = 0; i < splitedContent.length; i++) {
              let numberedString = (i + 1) + splitedContent[i];
              splitedContent.push(numberedString)
            }

            let modifiedFile: ModifiedFileModel = {
              name: x.name,
              splitedContent: splitedContent,
              testCoverage: x.testCoverage,
            };
            return modifiedFile;
          });
          console.log('modied--', this.modifiledFiles);
          this.currentFile = this.modifiledFiles[0];
        }
      })
    );
  }

  changeFile(file) {
    this.currentFile = file;
    // this.currentFile =
  }

  ngOnInit(): void {}
}
