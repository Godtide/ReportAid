#!/bin/sh

abigen --sol=./contracts/IATIActivities.sol --pkg=contracts --out=../server/internal/contracts/activities.go
abigen --sol=./contracts/IATIActivity.sol --pkg=contracts --out=../server/internal/contracts/activity.go
abigen --sol=./contracts/IATIActivityAdditional.sol --pkg=contracts --out=../server/internal/contracts/activityAdditional.go
abigen --sol=./contracts/IATIActivityBudgets.sol --pkg=contracts --out=../server/internal/contracts/activityBudgets.go
abigen --sol=./contracts/IATIActivityDates.sol --pkg=contracts --out=../server/internal/contracts/activityDates.go
abigen --sol=./contracts/IATIActivityParticipatingOrgs.sol --pkg=contracts --out=../server/internal/contracts/activityParticipatingOrgs.go
abigen --sol=./contracts/IATIActivityRelatedActivities.sol --pkg=contracts --out=../server/internal/contracts/activityRelatedActivities.go
abigen --sol=./contracts/IATIActivitySectors.sol --pkg=contracts --out=../server/internal/contracts/activitySectors.go
abigen --sol=./contracts/IATIActivityTerritories.sol --pkg=contracts --out=../server/internal/contracts/activityTerritories.go
abigen --sol=./contracts/IATIActivityTransactions.sol --pkg=contracts --out=../server/internal/contracts/activityTransactions.go
abigen --sol=./contracts/IATIBudgets.sol --pkg=contracts --out=../server/internal/contracts/budgets.go
abigen --sol=./contracts/IATIOrganisation.sol --pkg=contracts --out=../server/internal/contracts/organisation.go
abigen --sol=./contracts/IATIOrganisationBudgets.sol --pkg=contracts --out=../server/internal/contracts/organisationBudgets.go
abigen --sol=./contracts/IATIOrganisationCountryBudgets.sol --pkg=contracts --out=../server/internal/contracts/organisationCountryBudgets.go
abigen --sol=./contracts/IATIOrganisationDocs.sol --pkg=contracts --out=../server/internal/contracts/organisationDocs.go
abigen --sol=./contracts/IATIOrganisationExpenditure.sol --pkg=contracts --out=../server/internal/contracts/organisationExpenditure.go
abigen --sol=./contracts/IATIOrganisationRecipientBudgets.sol --pkg=contracts --out=../server/internal/contracts/organisationRecipientBudgets.go
abigen --sol=./contracts/IATIOrganisationRegionBudgets.sol --pkg=contracts --out=../server/internal/contracts/organisationRegionBudgets.go
abigen --sol=./contracts/IATIOrganisations.sol --pkg=contracts --out=../server/internal/contracts/organisations.go
abigen --sol=./contracts/IATIOrgs.sol --pkg=contracts --out=../server/internal/contracts/orgs.go
