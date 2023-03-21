
export function getRoA(drug) {
    switch (drug.roa) {
        case 1:
            return "Other"
        case 2:
            return "Oral"
        case 3:
            return "Buccal"
        case 4:
            return "Rectal"
        case 5:
            return "Inhaled"
        case 6:
            return "Sublabial"
        case 7:
            return "Intranasal"
        case 8:
            return "Sublingual"
        case 9:
            return "Injection"
        case 10:
            return "Buccal Injection"
        case 11:
            return "Intravenous Injection"
        case 12:
            return "Subcutaneous Injection"
        case 13:
            return "Intramuscular Injection"
        default:
            return ""
    }
}

export function getRoAJoined(drug) {
    const roa = getRoA(drug);
    if (roa) {
        return `, ${roa}`
    }

    return roa
}

export function getPrescribed(drug) {
    switch (drug.prescribed) {
        case 1:
            return "Over the counter"
        case 2:
            return "Prescribed by a doctor"
    }
}

export function getDose(drug) {
    if (drug.dosage === 0 && drug.dosage_unit === "") {
        return ""
    }

    if (drug.dosage === 0) {
        return drug.dosage_unit
    }

    let joiner = " "
    if (drug.dosage_unit.length < 5) {
        joiner = ""
    }

    return `${drug.dosage}${joiner}${drug.dosage_unit}`
}